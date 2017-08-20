package main

import (
	"time"
	"fmt"
)

// 计时器
// 在将来的某个时刻执行代码，或者以某种间隔重复执行。
// GO内置计时器和报警器功能使这两个任务变得容易。
func main() {
	timer1 := time.NewTimer(time.Second * 2)
	// 定时器将代表一个单一的事件。想等待多久，它提供
	// 一个通道，当时被通知。当前等待2秒。
	// 该<- time1.C计时器的通道块C，直到它发送指示所述定时器到期的值。
	<- timer1.C
	fmt.Println("Timer 1 expired")

	//等待， 可以使用time.Sleep。定时器可能有用的一个原因是
	//可以在定时器到期之前取消定时器。
	timer2 := time.NewTimer(time.Second)

	go func() {
		<- timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
		fmt.Println(stop2)
	}


}