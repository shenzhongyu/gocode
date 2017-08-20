package main

import "fmt"

// 非阻塞频道操作
// 通道上的基本发送和接收被阻止。我们可以使用
// select一个default字句实现非阻塞发送，接收，甚至阻塞
// 多路select。

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 这是一个非阻塞的接收。如果值可用，messages那么select将 <- message case
	// 使用该值。如果不是，将立即采取default情况
	select {
	case msg := <-messages:
		fmt.Println("recevied message", msg)
	default:
		fmt.Println("no message recevied")
	}
	// 非阻塞發送工作类似
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
		mg := <-messages
		fmt.Println(mg)
	default:
		fmt.Println("no message sent")
	}

	// 可以使用多个case以上的default子句实现多路非阻塞选择。这里我们尝试在两者message
	// 和非阻塞接收signals。
	select {
	case msg := <-messages:
		fmt.Println("recevied message", msg)
	case sig := <-signals:
		fmt.Println("recevied signal", sig)
	default:
		fmt.Println("no activity")
	}

}
