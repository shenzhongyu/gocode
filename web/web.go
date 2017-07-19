package main

import (
  "fmt"
  "strconv"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w,templateStr, r.URL.Path[1:], multiplicationTable())
}

func multiplicationTable() string{
  var str string
  for r := 1; r <= 9; r++ {
    str += "<tr>" 
    for c := 1; c <= r; c++ {
      str += "<td>"
      str += strconv.Itoa(c)
      str += " x "
      str += strconv.Itoa(r)
      str += " = "
      str += strconv.Itoa(r * c)
      str += "</td>"
    }
    str +="</tr>"
  } 
  return str
} 

func main() {
  http.HandleFunc("/",handler)
  http.ListenAndServe(":8090", nil)
}

const templateStr = `
<!DOCTYPE html>
<html lang="zh-CN">  
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>乘法口诀表</title>  
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <script src="https://cdn.bootcss.com/jquery/1.12.4/jquery.min.js"></script>
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    <style>
       html, body {
                background-color: #fff;
                color: #333;
                font-family: 'Raleway', sans-serif;
                font-weight: 100;
                height: 100vh;
                margin: 0;
            }
    </style>
 </head>  
  <body>  
    <div class="container-fluid">
      <div class="page-header">
	<h1>multiplication table <small>%s</small></h1>
      </div>
      <div class="table-responsive">
	<table class="table table-bordered table-hover center-block">%s</table>
      </div>
    </div>
  </body>  
</html>  
`