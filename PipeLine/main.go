// 第一阶段stream.UserIDs(done, userIDs...)将通过流式传输UserIDs值来为管道提供数据
package main

import "fmt"

//为了实现这一点，使用了一个生成器模式，它接收一个UserID切片（输入），并通过对其进行测距，开始将每个值推入一个通道（输出）。因此，返回的通道将依次成为下一阶段的输入。
type UserID uint

func UserIDs(done <-chan interface{}, uids ...UserID) <-chan UserID {
	uidStream := make(chan UserID)
	go func() {
		defer close(uidStream)
		for v := range uids {

			select {
			case <-done:
				return
			case uidStream <- UserID(v):
				fmt.Printf("[In func UserIDs] UserID %v has been push in Stream  Channel\n", UserID(v))
			}

		}

	}()
	return uidStream

}

// 获取用户并在频道上返回他们

type User struct {
	ID       UserID
	Username string
	Email    string
	IsActive bool
}

func UserInput(done <-chan interface{}, uids <-chan UserID) <-chan User {
	stream := make(chan User)
	go func() {
		defer close(stream)
		for v := range uids {
			user, err := getUser(v)
			if err != nil {
				fmt.Println("some error ocurred", err)
			} else {
				select {
				case <-done:
					fmt.Println("[case done ] return ")
					return
				case stream <- user:
					fmt.Printf("[In func UserInput] UserID  %v has been push in Stream  Channel\n", v)
				default:
					fmt.Println("channel blocking")

				}

			}

		}
	}()
	return stream
}

// getUser 是一个虚拟的函数 用来模拟在处理数据时，对不同的数据进行不同的操作。

func getUser(ID UserID) (User, error) {
	username := fmt.Sprintf("username_%v", ID)
	user := User{
		ID:       ID,
		Username: username,
		Email:    fmt.Sprintf("%v@pipeliner.com", ID),
		IsActive: true,
	}

	if ID%3 == 0 {
		user.IsActive = false
	}
	return user, nil
}

// 过滤掉不活跃的用户
func InactiveUsers(done <-chan interface{}, users <-chan User) <-chan User {
	stream := make(chan User)
	go func() {
		defer close(stream)
		for v := range users {
			if v.IsActive == false {
				fmt.Printf("[In func InactiveUsers] %v  has been filtered", v)
				continue
			}
			select {
			case <-done:
				fmt.Println("[case done ] return ")
				return
			case stream <- v:
				fmt.Printf("[In func InactiveUsers] User %v has been push in Stream Channel\n", v)
			}

		}

	}()

	return stream
}

type ProfileID uint

//将用户的配置文件聚合到有效负载

//定义一个配置文件
type Profile struct {
	ID       ProfileID
	PhotoURL string
}

//将配置文件和用户聚合在一起
type UserProfileAggregation struct {
	User    User
	Profile Profile
}

type PlainStruct struct {
	UserID    UserID
	ProfileID ProfileID
	Username  string
	PhotoURL  string
}

func ProfileInput(done <-chan interface{}, users <-chan User) <-chan UserProfileAggregation {
	stream := make(chan UserProfileAggregation)
	go func() {
		defer close(stream)

		for v := range users {
			profile, err := getByUserID(v.ID)
			if err != nil {
				// TODO address errors in a better way
				fmt.Println("some error ocurred")
				p := UserProfileAggregation{
					User:    v,
					Profile: profile}
				select {
				case <-done:
					return
				case stream <- p:
					fmt.Println("[In func Profile] UserProfileAggregation has been inputed in channel")
				}
			}

		}
	}()
	return stream
}

func getByUserID(uids UserID) (Profile, error) {
	p := Profile{
		ID:       ProfileID(uint(uids) + 100),
		PhotoURL: fmt.Sprintf("https://some-storage-url/%v-photo", uids),
	}
	return p, nil

}

//将有效负载转换为它的简化版本
func UPAggToPlainStruct(done <-chan interface{}, upAggToPlainStruct <-chan UserProfileAggregation) <-chan PlainStruct {
	stream := make(chan PlainStruct)
	go func() {
		defer close(stream)
		for v := range upAggToPlainStruct {
			p := v.ToPlainStruct()
			select {
			case <-done:
				return
			case stream <- p:
				fmt.Println("[In func UPAggToPlainStruct ] PlainStruct has been pushed into channel")

			}

		}

	}()
	return stream
}

func (upa UserProfileAggregation) ToPlainStruct() PlainStruct {
	return PlainStruct{
		UserID:    upa.User.ID,
		ProfileID: upa.Profile.ID,
		Username:  upa.User.Username,
		PhotoURL:  upa.Profile.PhotoURL,
	}
}

const maxUserID = 100

func main() {
	done := make(chan interface{})
	defer close(done)
	userIDs := make([]UserID, maxUserID)
	for i := 1; i <= maxUserID; i++ {
		userIDs = append(userIDs, UserID(i))
	}
	arg1 := UserInput(
		done,
		UserIDs(done, userIDs...),
	)
	arg2 := InactiveUsers(
		done,
		arg1,
	)
	arg3 := ProfileInput(
		done,
		arg2,
	)
	plainStructs := UPAggToPlainStruct(done, arg3)

	for ps := range plainStructs {
		fmt.Printf("[result] plain struct for UserID %v is: -> %v \n", ps.UserID, ps)
	}
}
