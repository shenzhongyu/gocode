package main

import (
	"fmt"
	"flag"
	"time"
	"github.com/shenzhongyu/gocode/demo_two/domain"
)

var infile *string = flag.String("i", "files/unsorted.dat", "File contains values for sorting.")
var outfile *string = flag.String("o", "files/sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")
var err1, err2, err3 error
var p = fmt.Println

func main() {
	flag.Parse()
	p("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)

	t1 := time.Now()
	todo := domain.NewFile(*infile, *outfile, *algorithm)
	err1 = todo.ReadValues()
	err2 = todo.SortValues()
	err3 = todo.WriteValues()
	t2 := time.Now()

	if err1 != nil || err2 != nil || err3 != nil {
		p(err1, "\n", err2, "\n", err3)
	} else {
		p("The sorting to book file process costs", t2.Sub(t1), "to complete.")
	}

}