package main

import "fmt"

// 浮点型数组声明
var a [5]float32

func main() {
	// 字符串型数组初始化
	b := [10]string{"J", "o", "k", "e", "t", "o", "r", "S", "e", "n"}

	fmt.Println("float32:", a)
	fmt.Println("string:", b)

	a[4] = 25.9

	fmt.Println("set:", a, "get:", a, "len:", len(a))

	arr := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("avg:", getAverage(arr, len(arr)))

	//二维数组声明d
	var c [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}

	fmt.Println("2d:", c)

	// 二(多)维数组初始化 var array_name [row][col] array_type
	d := [2][3]string{{"JoketorSen", "like", "apple"}, {"JoketorSen", "dislike", "origin"}}

	fmt.Println("2d:d", d, "get:", d[1][2])

}

// Go 语言向函数传递数组
func getAverage(arr [6]int, size int) float32 {
	var avg float32
	var sum int

	for i := 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum / size)

	return avg

}
