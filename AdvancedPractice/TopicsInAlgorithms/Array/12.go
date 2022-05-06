package main

import "fmt"

type Person interface {
	Create()
}

func NewPerson(name string) Person {
	if name == "student" {
		return Student{
			"2020213760",
			"xb",
		}

	}
	if name == "teacher" {
		return Teacher{
			3,
			"po",
			"3w",
		}
	}
	return nil
}

type Student struct {
	Id   string
	Name string
}

func (s Student) Create() {
	fmt.Println(s.Id, "has create")
}

type Teacher struct {
	Id    int
	Name  string
	Money string
}

func (t Teacher) Create() {
	fmt.Println(t.Id, "has create")
}

type Member struct {
	p   Person
	age int
}

func NewNumber() Member {
	return Member{
		NewPerson("student"),
		19,
	}
}

/* func (s Student) Create() {
	fmt.Println(s.name + "has create")
}

func (s Student) String() string {
	return fmt.Sprintf("(name:%v)\n(age:%v)", s.name, s.age)
}
func (s Student) Error() string {
	return fmt.Sprintf("Error-X: %s, %s error", s.name, s.age)
}
func NewStudent(name string, age string, score string) interface{} {
	return Student{name, age, score}
}
func NewStudent1(name string, age string, score string) Creater {
	return Student{name, age, score}
} */

func main() {
	p := NewNumber()
	fmt.Println(p.age)
	p.p.Create()

}
