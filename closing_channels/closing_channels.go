package main

import "fmt"

// 关闭渠道
// 关闭频道表示不会再发送更多的值。这可以用于将通道的
// 通道传送到通道的接收器。

func main() {
  jobs := make(chan int, 5)
  done := make(chan bool)
  
  // goroutines反复接收来自jobs用j,more ：= <- jobs
  // 再接收的这个特殊的二值形式中，该more值将是false
  // 如果jobs已经closed和在通道中的所有值都已经被接收到。
  // 使用它来通知done我们何时完成所有的工作。
  go func() {
    for {
      j, more := <- jobs
      if more {
	fmt.Println("received job", j)
      } else {
	fmt.Println("received all jobs")
	done <- true
	return
      }
    }
  }()
  
  // 通过jobs频道向工作人员发送3个作业，然后关闭它。
  for j := 1; j <= 3; j++ {
    jobs <- j
    fmt.Println("send job", j)
  }
  
  close(jobs)
  fmt.Println("send all jobs")
  
  // 我们等待着工人使用我们前面看到的同步(synchronization)方法。
  <- done
  
}
