package temp

import (
	"fmt"
	"regexp"
)

type User struct {
	//用户类别  true 为超级用户 其他为普通用户
	category           int
	detailsInformation []interface{}
}

func NewUser() User {
	u := User{}
	//u.detailsInformation = make([]interface{}, 10)

	fmt.Println("Please input user category:")
	fmt.Scan(&u.category)
	var temp string

	fmt.Println("Please input userStudentId:")
	fmt.Scan(&temp)
	found, err := regexp.MatchString("^[0-9]{10}$", temp)
	if err != nil {
		fmt.Printf("found=%v, err=%v", found, err)
	} else {
		u.detailsInformation = append(u.detailsInformation, temp)
	}
	fmt.Println("Please input userName:")
	fmt.Scan(&temp)
	found, err = regexp.MatchString("^[a-zA-Z0-9]$", temp)
	if err != nil {
		fmt.Printf("found=%v, err=%v", found, err)
	} else {
		u.detailsInformation = append(u.detailsInformation, temp)

	}
	return u
}
func (u *User) Search() {
	fmt.Println(u.detailsInformation...)
	fmt.Println("User Search Success!")
}
func main() {
	u := NewUser()
	u.Search()

}
