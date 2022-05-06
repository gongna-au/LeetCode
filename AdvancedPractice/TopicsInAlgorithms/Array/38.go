package main

import (
	"fmt"
	"time"
)

var Pending = "0"
var Resloved = "1"
var Rejected = "-1"

type handeler func(interface{})
type Task func(handeler, handeler)
type Then func(interface{}) interface{}
type lock struct {
	async chan int
	ret   interface{}
	state string
}
type Promise struct {
	*lock
}

func (l *lock) start() {
	close(l.async)
}
func (l *lock) end() {
	<-l.async
	l.state = Resloved
}

func (p *Promise) Await() interface{} {

	if p.state != Pending {
		return p.ret
	}
	<-p.async
	p.end()
	return p.ret
}
func (p *Promise) Then(callback Then) {
	p.Await()
	p.ret = callback(p.ret)

}
func exec(task Task) *Promise {
	async := &Promise{&lock{make(chan int), nil, Pending}}
	go async.Exec(task)
	return async
}
func (p *Promise) reslove(ret interface{}) {
	p.ret = ret
	p.state = Resloved
}
func (p *Promise) reject(ret interface{}) {
	p.ret = ret
	p.state = Rejected
}
func (p *Promise) Exec(task Task) {
	go p.end()
	task(p.reslove, p.reject)
	p.start()
}
func job(reslove, reject handeler) {
	fmt.Println("hi~")
	time.Sleep(time.Second * 2)
	fmt.Println("end")
	reslove(100)
}
func main() {
	result := exec(job)
	fmt.Println("bbb")
	fmt.Println(result.Await())
	fmt.Println("job done")
}
