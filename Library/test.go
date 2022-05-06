package main

import (
	"fmt"
)

type Open func(str string, url string)
type Onload func()
type Send func() int
type Response interface{}

type Model struct {
	open         Open
	onLoad       Onload
	resPonseType int
	send         Send
}

func NewModel() *Model {
	return &Model{
		Show,
		nil,
		0,
		SendActual,
	}
}

func Show(str string, url string) {
	fmt.Print(str + url)
}

func SendActual() int {
	fmt.Println("message has been send")
	return 200
}

type callbacktype func(int)

func Test(url string, c callbacktype) *Model {
	model := NewModel()
	model.open("Get", url)
	model.resPonseType = 0
	model.onLoad = func() {
		c(model.resPonseType)
	}
	model.send()
	return model
}

func callback(x int) {
	fmt.Println("I am main function", x)
}

func main() {

	p := Test("www.baidu", func(x int) {
		fmt.Println("I am main function", x)
	})
	p.onLoad()

}
