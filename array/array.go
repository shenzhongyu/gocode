 package main
 
 import "fmt"
 
 func main() {
   // 一维数组
   var arr1 = [5]int{1, 2, 3, 4, 5}
   // 多维数组 [row][col] typ
   var arr2 = [2][3]string{{"joketorsen", "like", "fruit"},{"kitty", "like", "monkey"}}
   fmt.Println(arr1, arr2, getAverage(arr1, 4))
}

// 向函数传递数组
func  getAverage(arr [5]int, size int) float32 {
   var i, sum int
   var avg float32
   for i = 0; i < size; i++ {
      sum += arr[1]
   }
     
  avg = float32(sum / size)
     
  return avg
     
}  
