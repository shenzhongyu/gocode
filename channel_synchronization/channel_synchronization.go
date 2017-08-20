package main

import (
	"fmt"
	"time"
)

// 通道同步
// 使用渠道来同步执行goroutines。使用阻止接收等待goroutines完成。

// 该done通道将用于通知另一个goroutine该功能的工作已完成。
// 发送一个值以通知我们已经完成。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println(time.Second)
	fmt.Println("done")

	done <- true

}

func main() {

	done := make(chan bool, 1)
	go worker(done)
	//阻止，直到我们受到通道上的工作人员的通知。

	// <- done从该程序中删除了该行，程序将在worker甚至启动之前退出。
	<-done

}
