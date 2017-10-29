package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"strconv"
	"time"
	"github.com/shenzhongyu/gocode/demo_one/algorithm_sort"
	"io"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort alogorithm")
var p = fmt.Println

func main() {
	flag.Parse()

	if infile != nil {
		p("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)

	}

	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
			case "qsort":
				algorithm_sort.QuickSort(values)
			case "bubblesort":
				algorithm_sort.BubbleSort(values)
			default:
				p("Sorting algorithm", *algorithm, "is either unknow or unsupported.")
		}
		t2 := time.Now()

		p("The sorting process costs", t2.Sub(t1), "to complete.")

		writeValues(values, *outfile)
	} else {
		p(err)
	}

}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err !=nil {
		p("Failed to open the input file", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			p("A too long line, seens unexpected.")
			return
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return

}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		p("Failed to created the output file", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}


