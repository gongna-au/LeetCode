package main

import (
	"errors"
	"fmt"
)

/*
实现一个计算功能：一个数从0开始，每次加上自己的值和当前循环次数（当前第几次，循环从0开始，到9，共10次），然后*2，这样迭代10次：

*/
type Creater interface {
	Create()
}

func NewCreater(args string) (Creater, error) {
	if args == "student" {
		return &Student{}, nil
	}
	if args == "record" {
		return &Record{}, nil
	}
	return nil, errors.New("NewCreater error")

}

type Student struct {
}

func (s *Student) Create() {
	fmt.Println("student has create")

}

type Record struct {
}

func (r *Record) Create() {
	fmt.Println("record has create")

}

func main() {
	//var temp1 interface{}
	var temp2 interface{}

	s1 := &Student{}
	temp2 = s1
	switch v := temp2.(type) {
	case Creater:
		fmt.Println("is Creater interface")

	case *Student:
		v.Create()
	case *Record:
		v.Create()

	}

}
