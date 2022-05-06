package main

import "fmt"

type middleware func(hander) hander
type hander func(int, int)

func Sub(a int, b int) {
	fmt.Println("InSub a=", a)
	fmt.Println("InSub b=", b)
	fmt.Println(a - b)
}
func Add(a int, b int) {
	fmt.Println("InAdd a=", a)
	fmt.Println("InAdd b=", b)
	fmt.Println(a + b)
}
func Complete(p func(int, int)) func(int, int) {
	return func(a int, b int) {
		fmt.Println("a*b=", a*b)
		p(a, b)
	}
}
func Logging() middleware {
	return func(f hander) hander {

		return func(a int, b int) {
			f(a, b)

		}
	}
}
func Chain(f hander, mid1 middleware, mid2 middleware) (h1 hander, h2 hander) {
	f1 := mid1(f)
	f2 := mid2(f)

	return f1, f2
}

func main() {
	middle1 := Logging()
	middle2 := Logging()
	result1, result2 := Chain(Sub, middle1, middle2)
	result1(1, 2)
	result2(3, 4)
}
