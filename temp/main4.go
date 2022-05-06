package main

import (
	"fmt"
	"regexp"
)

type App struct {
	menu MenuService
}

func NewApp() *App {
	s := NewSecondaryMenu()
	return &App{
		s,
	}

}
func (a *App) Run() {

	switch k := a.menu.(type) {
	case *Menu:
		k.Listen()
		k.Execute()
	}
}

type MenuService interface {
	Listen()
	Execute()
}

//MenuService级别
type Menu struct {
	item Item
}
type Item interface {
	Method()
}

func NewSecondaryMenu() *Menu {

	e := NewEventToExecute()
	return &Menu{
		e,
	}
}
func (s *Menu) Listen() {
	fmt.Println("Menu Start Listening")
	switch k := s.item.(type) {
	case *EventToExecute:

		k.GetEvent()

	case *User:
		k.InputInformation()

	}
}
func (s *Menu) Execute() {
	fmt.Println("Menu Execute start")
	switch k := s.item.(type) {
	case *User:
		k.Method()
	}
	fmt.Printf("Menu Execute end")
}

type EventToExecute struct {
	event Event
}
type Event interface {
}

func NewEventToExecute() *EventToExecute {
	return &EventToExecute{
		nil,
	}
}

func (e *EventToExecute) Method() {
	button := newButton("\n>>1.新增用户\n>>2.查找用户\n>>3.用户登陆\n>>4.用户注销\n>>5.退出")
	button.showPageInformation()
}

func (e *EventToExecute) GetEvent() {
	button := newButton("\n>>1.新增用户\n>>2.查找用户\n>>3.用户登陆\n>>4.用户注销\n>>5.退出")
	button.showPageInformation()

	button.listenButtonBePressed()

}

type Button struct {
	information string
	bePressed   int
}

func newButton(info string) *Button {
	return &Button{
		info,
		0,
	}
}
func (b *Button) showPageInformation() {
	fmt.Println(b.information)
}

type handler func()

func (b *Button) listenButtonBePressed(h handler) {
	fmt.Println("listen button be pressed.....")
	fmt.Scan(&b.bePressed)

}

//Item级别
type User struct {
	//用户类别  true 为超级用户 其他为普通用户
	category           int
	detailsInformation []interface{}
}

func NewUser() User {
	u := User{}
	return u
}

func (u *User) Method() {

	button := newButton("\n>>1.新增用户\n>>2.查找用户\n>>3.用户登陆\n>>4.用户注销\n>>5.退出")
	button.showPageInformation()
	switch button.bePressed {
	case 1:
		{
			u.Insert()
		}
	case 2:
		{
			u.Search()
		}
	case 3:
		{
			u.Login()
		}
	case 4:
		{
			u.Delete()
		}
	}

}

func (u *User) InputInformation() {
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
}

func (u *User) Insert() {
	fmt.Println(u.detailsInformation...)
	fmt.Println("User Insert Success!")
}
func (u *User) Search() {
	fmt.Println(u.detailsInformation...)
	fmt.Println("User Search Success!")
}
func (u *User) Login() {
	fmt.Println("User Login Success!")
}
func (u *User) Delete() {
	fmt.Println("User Delete Success!")
}

func main() {
	app := NewApp()
	app.Run()

}
