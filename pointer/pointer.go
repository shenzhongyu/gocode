package main

import "fmt"

var (
  a *int
  arr [3]*int
)

func main() {
  
  var b int = 10
  fmt.Printf("变量的地址：%x\n", &a)
  fmt.Printf("变量的地址：%x\n", &b)
  
  // a 变量实际值为0
  if a != nil {
    fmt.Printf("a 不是空指针\n")
  } else {
    fmt.Printf("a 是空指针\n")
  }
  
  // 指针数组
  c := []int{10, 100, 200}
  for i := 0; i < 3; i ++ {
    arr[i] = &c[i] // 整数地址赋值给指针数组
  }
  
  for j := 0; j < 3; j ++ {
      fmt.Printf("a[%d] = %d\n", j,*arr[j] )
   }
  
}