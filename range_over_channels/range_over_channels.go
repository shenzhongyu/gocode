package main

import "fmt"

//通道上的范围
// 在前面的例子中，for并range提供了基本数据结构的迭代。
// 还可以使用这种语法来迭代从通道接收的值。
func main() {
	queue := make(chan string, 2)
	//在队列queue频道中迭代超过2个值。
	queue <- "one"
	queue <- "two"
	close(queue)
	// range会从它接收到的每个元素进行迭代queue。
	// 因为我们close是上面的通道，所以迭代在接收到
	// 2个元素之后终止。
	for elem := range queue {
		fmt.Println(elem)
	}
	// 该示例还显示可以关闭非空通道，但任是可以接收剩余的值。
}
