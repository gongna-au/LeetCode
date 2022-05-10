package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var num int

var (
	// 新用户到来，通过该 channel 进行登记
	userEnterChannel = make(chan *User)
	// 用户离开，通过哎 channel 进行登记
	userLeavingChannel = make(chan *User)
	// 广播专用的用户普通信息 channel，缓冲是尽可能避免出现异常情况堵塞，这里简单给了 8，
	// 具体值根据情况调整
	messageChannel = make(chan string, 8)
)

type User struct {
	//每个用户的唯一标识
	ID int
	//每个用户的ip地址和端口
	Addr string
	//每个用户进入聊天室的时间
	EnterTime time.Time
	//每个用户有一个专属于自己接受信息的通道
	//当需要给某个用户发送消息时，只需在该用户的 MessageChannel 中写入消息即可。
	MessageChannel chan string
}

//四种 goroutine: main goroutine
//广播消息 goroutine
//每一个客户端链接都会有一个读和写的 goroutine
func main() {
	//在 listen 时没有指定 IP 地址，所以表示绑定到当前机器的 IP 地址上
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		panic(err)
	}
	go broadCaseter()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		} else {
			go handleConn(conn)

		}

	}

}

// broadcaster 用于记录聊天室用户，并进行消息广播
// 1. 新用户进来 2. 用户普通消息 3. 用户离开
func broadCaseter() {
	//用户所有的用户是一个map
	users := make(map[*User]struct{})
	//在登记/注销用户时，是通过 map 存储在线用户的
	//在用户登记、注销时应使用专门的 channel。
	//在注销时，除从 map 中删除用户外，还会关闭 user 的 MessageChannel，避免出现 goroutine 泄漏问题
	//全局的 messageChannle 一般用来给聊天室中的所有用户广播消息
	//broadcaster 函数的关键在于对 goroutine 和 channel 的运用。
	//它很好地践行了 Go 的概念，即通过通信来共享内存。它里面的三个 channel 的定义如下:
	for {
		select {
		case user := <-userEnterChannel:
			// 新用户进入
			users[user] = struct{}{}
		case user := <-userLeavingChannel:
			// 用户离开
			delete(users, user)
			//避免 goroutine 泄漏
			close(user.MessageChannel)
		case msg := <-messageChannel:
			for user := range users {
				user.MessageChannel <- msg
			}
		}
	}

}
func GetUserId() int {
	num = num + 1
	return num

}

func handleConn(conn net.Conn) {
	//1.处理完毕后关闭连接
	defer conn.Close()
	//2.每个连接进入到这个函数里面就保存我们需要的数据
	user := &User{
		GetUserId(),
		conn.RemoteAddr().String(),
		time.Now(),
		make(chan string, 8),
	}
	//go sendMessage(conn, user.MessageChannel)
	//也就是在往user.MessageChannel这个里面写的时候就是用户可以接收到外界的信息
	//那么还需要实现用户可以往外发送信息
	go sendMessageOut(conn, user.MessageChannel)

	//3.给当前用户发送欢迎信息 往每个进入聊天室的用户的Channel里面写
	//给当前用户发送欢迎信息，同时给聊天室所有用户发送有新用户到来的提醒
	//注意顺序一定是先欢迎用户，然后才把用户添加到全局聊天室的用户列表当中
	//否则会收到自己到来的消息提示
	//当然，也可以做消息过滤处理
	user.MessageChannel <- "Welcome " + strconv.Itoa(user.ID)
	//4. 需要一个公共的messageChannel让所有的用户知道这个用户来了
	messageChannel <- "User:" + strconv.Itoa(user.ID) + "has enter"
	//5. 把每个进入的用户的指针添加到全局的变量中
	//将该用户写入全局用户列表，即聊天室用户列表
	userEnterChannel <- user
	//6 .循环的读取每个用户的输入
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messageChannel <- strconv.Itoa(user.ID) + ":" + input.Text()
	}
	if err := input.Err(); err != nil {
		log.Println("从用户读取数据错误！", err)
	}
	//6.用户离开
	userLeavingChannel <- user
	messageChannel <- "User:" + strconv.Itoa(user.ID) + "has left"

}

/*
需要注意的是，因为 sendMessage 在一个新的 goroutine 中，所以如果函数里的 ch 不关闭
则该 goroutine 是不会退出的，因此需要注意不关闭 ch 导致的 goroutine 泄漏的问题这里有一个语法
ch <-chan string
channel 实际上有三种类型，大部分时候，我们只用了其中一种，即正常的既能发送也能接收的 channel。
除此之外，还有单向的 channel： 只能接收和只能发送。它们无法直接被创建，而是通过正常(双向) channel 转换而来(会自动隐式转换)。
比如在上面代码中的 ch <-chan string 就是为了限制在 sendMessage 函数只从 channel 中读取数据，不允许往里写数据 */
func sendMessageOut(conn net.Conn, ch <-chan string) {
	//向外界发送信息已经存储在自己的chan string里面
	//现在可以直接读取出来
	for msg := range ch {
		//把读到的数据又写入到了连接conn
		fmt.Fprintln(conn, msg)
	}

}
