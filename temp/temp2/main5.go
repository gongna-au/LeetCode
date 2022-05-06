package main

import (
	"fmt"
	"regexp"
	//"github.com/FreshmanGuidanceProject-backend/api/router"
)

type Button struct {
	information string
	bePressed   int
	handler     []handler
}

func newButton(info string) *Button {
	return &Button{
		info,
		0,
		nil,
	}
}
func (b *Button) showPageInformation() {
	fmt.Println(b.information)
}

type handler func()

func (b *Button) loadHandler(h []handler) {
	for _, v := range h {
		b.handler = append(b.handler, v)
	}
}
func (b *Button) listenButtonBePressed() {
	fmt.Println("listen button be pressed.....")
	fmt.Scan(&b.bePressed)
	fmt.Println("button has been pressed.....")
	b.handler[b.bePressed-1]()
}

type Book struct {
	category           string
	detailsInformation map[string]interface{}
}

type User struct {
	//用户类别  true 为超级用户 其他为普通用户
	category           int
	detailsInformation map[string]interface{}
}

func NewUser() User {
	u := User{
		0,
		make(map[string]interface{}),
	}
	return u
}

func (u *User) InputInformation(strs []string) error {
	for _, v := range strs {
		fmt.Println("Please input" + " " + v)
		var temp string
		fmt.Scan(&temp)
		found, err := regexp.MatchString("^[0-9]{10}$", temp)
		if err != nil {
			fmt.Printf("found=%v, err=%v", found, err)
			return err
		} else {
			u.detailsInformation[v] = temp
		}
	}
	return nil

}
func (u *User) VerifyLogin() {
	strs := []string{"userId", "userPassword"}
	err := u.InputInformation(strs)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	//数据查询
	SearchInDB(u)
	fmt.Println("Verify Login Success!")

}

type SearchInDBService interface {
	Search()
}

func SearchInDB(s SearchInDBService) {
	switch k := s.(type) {
	case *User:
		k.Search()
	}
}

func (u *User) Insert() {
	for _, v := range u.detailsInformation {
		fmt.Println(v)
	}
	fmt.Println("User Insert Success!")
}
func (u *User) Search() {
	for v := range u.detailsInformation {
		fmt.Println(v)
	}
	fmt.Println("User Search Success!")
}
func (u *User) Login() {
	fmt.Println("User Login Success!")
}
func (u *User) Delete() {
	fmt.Println("User Delete Success!")
}

func handler11() {
	fmt.Println("Borrow book Success!")
}
func handler12() {
	fmt.Println("Back book Success!")
}
func handler13() {
	fmt.Println("Search book Success!")
}

func handler1() {
	fmt.Println("Hander1 Start!")
	button := newButton("\n>>1.借阅书籍\n>>2.归还书籍\n>>3.查询书籍\n>>4.返回\n>>5.退出")
	button.showPageInformation()
	array := []handler{
		handler11,
		handler12,
		handler13,
	}
	button.loadHandler(array)
	button.listenButtonBePressed()
	fmt.Println("Hander1 Success!")

}
func handler2() {
	fmt.Println("Hander2 Start!")
	u := NewUser()
	strs := []string{"userId", "userPassword"}
	err := u.InputInformation(strs)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	u.Insert()
	fmt.Println("Hander2 Success!")

}
func handler3() {
	fmt.Println("Hander3 Start!")
	u := NewUser()
	strs := []string{"userId", "userPassword"}
	err := u.InputInformation(strs)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	u.Login()
	fmt.Println("Hander3 Success!")
}
func handler4() {
	fmt.Println("Hander4 Start!")
	u := NewUser()
	strs := []string{"userId", "userPassword"}
	err := u.InputInformation(strs)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	u.category = 1
	u.Insert()
	fmt.Println("Hander4 Success!")
}
func handler5() {
	fmt.Println("Hander5 Start!")

	fmt.Println("Hander5 Success!")

}

func main() {
	button := newButton("\n>>1.读者登陆\n>>2.读者注册\n>>3.管理员登陆\n>>4.管理员注册\n>>5.查询书籍\n>>6.退出")
	button.showPageInformation()
	array := []handler{
		handler1,
		handler2,
	}
	button.loadHandler(array)
	button.listenButtonBePressed()

}
