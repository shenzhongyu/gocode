package main

import "fmt"

const(
  UNKNOWN iota
  MALE
  FAMALE
)

var sex string

func main() {

  if UNKNOWN == 0 {
    sex = "阴阳人" 
  } else if MALE == 1 {
    sex = "男人"
  } else if FAMALE == 2 {
    sex = "女人"
  }
  
  fmt.Printf(sex)

}
