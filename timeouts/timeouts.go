package main

import (
	"fmt"
	"time"
)

// 超时
// 超时对于连接到外部资源的程序或其他需要绑定执行时间
// 的程序很重要。在GO中实现超时很容易和优雅，感谢渠道
// 和slect.

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// select 实现超时。由于select第一次收到的收益
	// 已经准备就绪，如果操作超过允许的1s，我们
	// 将采取超时的情况。
	select {
	// c1在2s后返回一个通道的结果
	case res := <-c1:
		fmt.Println(res)
	// <- time.After等待1s超时后发送一个值
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
