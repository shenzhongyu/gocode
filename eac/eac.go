package main

import "fmt"

func main() {
   i := 1
   for i <= 3 {
     fmt.Println(i)
     i = i + 1
   }
   
   for j := 2; j <= 10; j++ {
     fmt.Printf("这是第 %d 行\n",j)
   }   
  
   for {
     fmt.Printf("loop")
     break
    
   } 

   for n := 0; n < 20; n++ {
   
      if(n % 2 == 0){
     
         continue
	
       }

       fmt.Printf("打印第 %d 行的数字 \n", n)

   }
   
   //乘法口诀表
   for  c := 1; c <= 9; c++ {
     for r := 1; r <= c; r++ {
       fmt.Println("|", r, "x", c ,"=", r*c)
    }
  }


}

