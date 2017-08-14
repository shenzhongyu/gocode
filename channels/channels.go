package main 

import "fmt"

type user struct{
    name, sex string
    age int
}

// 通道是连接并发goroutines的管道。我们可以从一个goroutines发送值
// 到通道，并将这些值接收到另一个goroutines中。

func main() {
  // 创建一个新的频道make(chan val-type)。渠道是由他们传达的
  // 价值键入。使用channel <- 语法将值发送到通道
  message := make(chan string)
  user_info := make(chan user)
  
  go func() {
    message <- "ping"
    user_info <- user{ name:"JoketorSen", age:17, sex:"男" }
  }()
  
  // "ping"的消息通过我们的通道成功地从一个goroutine传递到另一个
  // groutine。
  msg := <- message
  user := <- user_info
  
  fmt.Printf("接收通道:%s\n", msg)
  fmt.Println(user, user.name)
  
}
