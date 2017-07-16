 /*运算符篇*/
package main

import "fmt"

var a, b = 20, 10
var (
  c int
  d int32
  e float32
  ptr *int // 指针变量
  r string
)

// 算术运算符
func Arithmetic(operator string) int {
  switch operator {
    case "+" : 
      c = a + b
    case "-" :
      c = a - b
    case "*" :
      c = a * b
    case "/" :
      c = a / b
    case "%" :
      c = a % b
    case "++" :
      a ++ ; c = a
    case "--" :
      a -- ; c = a
    default : 
      c = 0
  }
  return c
}
// 关系运算符
func Relation(operator string) string {
  switch operator {
    case "==" :
      if(a == b) {
	 r = " 第一行 - a 等于 b\n"
      } else {
	 r = " 第一行 - a 不等于 b\n"
      }
    case "<" :
      if(a < b) {
	r = "第二行 - a 小于 b\n"
      } else {
	r = "第二行 - a 不小于 b\n"
      }
    case ">" :
      if(a > b) {
	r = "第三行 - a 大于 b\n"
      } else {
	r = "第三行 - a 不大于 b\n" 
      }
    case ">=" :
      if(a >= b) {
	r = "第四行 - a大于等于 b\n"
      }
    case "<=" :
      if(a <= b) {
	r = "第五行 - a 小于等于  b\n"
      } else {
	r = "第五行 - a 不小于等于  b\n"
      }  
    case "!=" :
      if(a != b){
	r = "第六行 - a 不等于 b\n"
      } 
    default :
        r =  ""
  }
  return r
}
// 逻辑运算符
func Logic(operator string) string {
  const T, F = true, false
  switch operator {
    case "&&" :
      if(T && F) {
	//
      } else {
	r = " 第一行 - T && F 为 false\n"
      }
    case "||" :
      if(T || F) {
	r = "第二行 - T || F 为 true\n"
      }
    case "!" : 
      if(! (T && F)) {
	r =  "第三行 - ! (T && F) 为 true\n"
      }
    default :
      r = ""  
  }
  return r
}
// 位运算符
func Post(operator string) int {
  switch operator {
    case "&" :
      c = a & b
    case "|" :
      c = a | b
    case "^" :
      c = a ^ b
    case "《" :
      c = a >> 2
    case "》" :
      c = a << 2
    default :
      c = 0
  }
  return c
}
// 赋值运算符
func Evaluate(operator string) int{
  a, c = 2, 10
  switch operator {
    case "=" :
      c = a
    case "+=" :
      c += a
    case "-=" :
      c -= a
    case "*=" :
      c *= a
    case "/=" :
      c /= a
    case "%=" :
      c %= a
    case "《=" :
      c <<= 2
    case "》=" :
      c <<= 2
    case "&=" :
      c &= a
    case "^=" :
      c ^= a
    case "|=" :
      c |= a
    default :
      c = a
  }
  return c
}

func main() {
  // 其他运算符
  ptr = &a // 'ptr' 包含了 'a' 	变量实际存储地址
  fmt.Println("c 的值为:",c ,"d变量类型为:",d , "d变量类型为:",e, "a变量类型为:",a, "ptr变量类型为:",ptr)
  // 算术运算符
  fmt.Println(Arithmetic("+"), Arithmetic("-"), Arithmetic("*"), Arithmetic("/"), Arithmetic("%"), Arithmetic("++"), Arithmetic("--"))
  // 关系运算符
  fmt.Println(Relation("=="), Relation("<"), Relation(">"), Relation(">="), Relation("<="), Relation("!="))
  // 逻辑运算符
  fmt.Println(Logic("&&"), Logic("||"), Logic("!"))
  // 位运算符
  fmt.Println(Post("&"), Post("|"),  Post("^"), Post("》"), Post("《"))
  // 赋值运算符
  fmt.Println(Evaluate("="), Evaluate("+="),  Evaluate("-="), Evaluate("*="), Evaluate("/="), Evaluate("%="), Evaluate("《="), Evaluate("》="), Evaluate("&="), Evaluate("^="), Evaluate("|="))
}