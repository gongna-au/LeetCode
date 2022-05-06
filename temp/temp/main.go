package main

import "fmt"

type Actions interface {
	execute()

}
func (m *Menu)execute(){
	

}

type Menu struct {
	menuText  string
	items    []*MenuItem
}
func NewMenu(text string,) Menu{
	return Menu{
		menuText :text,


	}


}
func (m *Menu) ShowMenu(){


}

type ActionProcess func()
type MenuItem struct {
	text  string
	callBack ActionProcess
}


type Application struct {
	actions []*Actions
}

func (a *Application) Run() {
	for _,v := range a.actions{
		if v!=nil{
			v.

		}


	}

}
func (a *Application) pushAction(newAction *Actions) {
	a.actions = append(a.actions, newAction)

}
