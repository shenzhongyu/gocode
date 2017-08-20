package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg []int) {
		fmt.Println("going:", msg)
	}([]int{1, 2, 3})

	var input int
	fmt.Println("打印变量指针:", &input)
	fmt.Scanln(&input)
	age := input
	fmt.Println(age * 2)
	fmt.Println("done")

}
