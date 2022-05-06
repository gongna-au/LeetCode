package main

import (
	"fmt"
	"sync"
)

type Value struct {
	mu  sync.Mutex
	val int
}

func (v *Value) Incr() {
	defer v.mu.Unlock()

	v.mu.Lock()
	v.val++
}

func main() {
	var i Value

	i.Incr()
	fmt.Println(i.val)

}
