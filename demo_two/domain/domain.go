package domain

import (
	"os"
	"errors"
	"bufio"
	"io"
	"strconv"
	"github.com/shenzhongyu/gocode/demo_two/algorithm"
)

type File struct {
	InFilePath, OutFilePath, algorithm string
	Values []int
}

// 初始化
func NewFile(inFile, outFile, algorithm string) *File {
	return &File{inFile, outFile, algorithm, make([]int, 0)}
}

// 读取文件内部数据
func (f *File) ReadValues() (err error) {
	file, err := os.Open(f.InFilePath)
	if err != nil {
		err = errors.New("Failed to open the input file." + f.InFilePath)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	f.Values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			err = errors.New("A too lang, seems unexpected.")
			return
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		f.Values = append(f.Values, value)
	}

	return
}

// 文件内部值作排序处理
func (f *File) SortValues() (err error) {
	if len(f.Values) == 0 {
		err = errors.New("File values were null.")
		return
	}
	switch f.algorithm {
	case "qsort":
		algorithm.QuickSort(f.Values)
		break
	case "bubblesort":
		algorithm.BubbleSort(f.Values)
	default:
		err = errors.New("Sorting a algorithm " + f.algorithm + "is either unknow or unsupported.")
	}
	return
}

// 排序后值写入文件方法
func (f *File) WriteValues() (err error) {
	if len(f.Values) == 0 {
		err = errors.New("File values were null.")
		return
	}

	file, err := os.Create(f.OutFilePath)
	if err != nil {
		err = errors.New("Failed to create output file." + f.OutFilePath)
		return
	}

	defer file.Close()

	for _, value := range f.Values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return
}