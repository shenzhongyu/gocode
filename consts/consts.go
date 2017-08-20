/*常量篇*/
package main

import "fmt"

const (
	PI      float32 = 3.1415926535 // 圆周率 浮点型
	C0      uint32  = 300000000    // 真空光速 整数型 单位m/s
	VERSION string  = "v~1.0.0"    // 版本号 字符串型
)

const (
	UNKNOWN = iota // 未知 0
	MALE           // 男 1
	FAMALE         // 女 2
)

const (
	ONE   = iota         // const 第一次定义 iota = 0
	TWO                  // 第二次自增1     iota += 1
	THREE = "JoketorSen" // 类型改变，无法实行算术运算，常规常量定义
	FOUR                 // iota范围，未被赋值，默认承接上一次值
	FIVE  = 80           // 类型复原，常量被赋值80,常规常量定义
	SIX                  // iota范围，未被赋值，默认承接上一次值
	SEVEN = iota         // 第七次自增1 恢复计数 6
	EIGHT                // 第八次自增1 恢复计数 7

)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {

	fmt.Println("圆周率：", PI, ";真空光速：", C0, ";版本号：", VERSION)
	fmt.Println("未知：", UNKNOWN, ";男：", MALE, ";女：", FAMALE)
	fmt.Println("ONE：", ONE, ";TWO：", TWO, ";THREE：", THREE, ";FOUR:", FOUR, ";FIVE:", FIVE, ";SIX:", SIX, ";SEVEN:", SEVEN, ";EIGHT:", EIGHT)
	fmt.Println("i = ", i, ";j = ", j, ";k = ", k, ";l = ", l)

}
