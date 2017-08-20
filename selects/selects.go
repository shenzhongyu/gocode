package main

import (
	"fmt"
	"time"
)

// 选择
// 选择可以让我们等待多个频道操作。将goroutines和频道与精选组合是GO的强大功能。

func main() {
	// 声明两个渠道
	c1 := make(chan string)
	c2 := make(chan string)

	// goroutines通道c1存值
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	// goroutines通道c2存值
	go func() {
		// 每个通道将在一段时间后接收一个值，以模拟例如阻止在
		// 并行的goroutines中执行的RPC操作。
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// 选择两个渠道，select同时等待两个值，最后打印。
		// 总执行时间只有～2s ，c1和c2同时Sleeps同时执行。
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

}
