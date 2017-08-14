package main

import "fmt"

// 默认情况下通道是无缓冲的，这意味着chan <- 如果有一个
// 相应的receive(<- chan)准备好接收发送的值， 它们将只接受
// sent()。缓冲通道接受有限数量的值，而没有这些值的相应接收器。
func main() {
  // make有一个缓冲最多2个值的字符串通道。
  messages := make(chan string, 2)
  fmt.Println(messages)
  
  messages <- "buffered"
  messages <- "channel"
  
  // 应为这个通道是缓冲的，所以我们可以将这些值发送到通道中，
  // 而没有相应的并发接收。
  fmt.Println("缓冲：", <- messages)
  fmt.Println("通道：", <- messages)
  
}