package main

import (
	"fmt"
	//"time"
)

type Orlder struct {
	id int
}

func (orlder *Orlder) String() string {
	return fmt.Sprintln("Orlder ", orlder.id, " has finish")
}
func ProducerOrlder(out chan<- Orlder) {
	for i := 0; i < 20; i++ {
		temp := Orlder{
			i,
		}
		out <- temp

	}
	close(out)

}
func ConsumerOrlder(in <-chan Orlder) {
	for {
		k, ok := <-in
		if ok == true {
			//处理每一笔订单的逻辑
			fmt.Println(k)
		} else {
			break
		}
	}

}

func main() {
	ch := make(chan Orlder)
	go ProducerOrlder(ch)
	ConsumerOrlder(ch)
	/* time.Sleep(time.Second * 1) */

}
