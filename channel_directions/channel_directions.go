package main

import "fmt"

type user struct{
  name, sex string
  age int
}
// 渠道路线
// 使用频道作为功能参数时，您可以指定频道是仅仅发送或者接收值。
// 这种特异性增加了程序的类型安全性。

// ping 功能仅接收用于发送值的通道。
// 尝试在此频道上收到编译时错误。
func ping(pings chan <- user, msg user) {
  pings <- msg
}

// pong函数接收一个通道拥用于receive(pings), 另一个用于sending(pongs)
func pong(pings <- chan user, pongs chan <- user) {
  msg := <- pings
  pongs <- msg
}

func main() {
  pings := make(chan user, 1)
  pongs := make(chan user, 1)
  ping(pings, user{name:"JoketorSen", sex: "男", age:  18 })
  pong(pings, pongs)
  fmt.Println(<- pongs)
}
