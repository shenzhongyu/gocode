package main

import (
	"fmt"
	"os"
)

var p = fmt.Println

func main() {
	p("hello world!")

	beyondHello()
}

func beyondHello() {
	var x int
	x = 3
	y := 4
	sum, prod := learnMultiple(x, y)
	p("sum:", sum, "prod:", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	str := "Learn Go!"
	s2 := `A "raw" string literal can include line breaks.`
	p(str, s2)

	g := 'Σ'
	f := 3.14159
	c := 3 + 4i
	p(g, f, c)

	var u uint = 7
	var pi float32 = 22. / 7
	n := byte('\n')
	p(u, pi, n)

	var a4 [4]int
	a5 := [...]int{3, 1, 4, 57, 100}
	p(a4, a5)

	s3 := []int{1, 4, 7}
	s4 := make([]int, 5)
	var d2 [][]float64
	bs := []byte("a slice")
	p(s3, s4, d2, bs)

	s := []int{1, 2, 3}
	s = append(s, 2, 5, 65)
	p(s)

	s = append(s, []int{7, 8, 9}...)
	p(s)

	j, q := learnMemory()
	p(*j, *q)

	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1
	p(m)

	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "This is How you write to a file, by the way.")
	file.Close()

	learnFlowControl()
}


func learnMemory() (j, q *int) {
	j = new(int)
	s := make([]int, 20)
	s[3] = 7
	r := -2
	return &s[3], &r
}

func learnFlowControl () {
	if true {
		p("told ya")
	}

	if false {
		p("Pout")
	} else {
		p("Gloat")
	}

	x := 42.0
	switch x {
	case 0:
	case 1:
	case 42:
	case 43:
	default:
	}

	for x := 0; x < 3; x++ {
		p("iteration", x)
	}

	for {
		break
		continue
	}

	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}

	for _,name := range[]string{"Bob", "Bill", "Joe"}{
		fmt.Printf("Hello, %s\n", name)
	}

}