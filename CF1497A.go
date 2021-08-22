package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1497A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1497A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1497A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1497A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1497A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1497A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1497A
 Date: 21/08/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1497/A
 Problem: CF1497A
**/
func (in *CF1497A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		arr := []int{}
		for i := 0; i < n; i++ {
			num := in.NextInt()
			arr = append(arr, num)
		}
		sort.Ints(arr)
		last := []int{}
		index := 0
		for i := 0; i < n-1; i++ {
			if arr[i] == arr[i+1] {
				num := arr[i]
				last = append(last, num)
				arr[i] = -1
				index++
			}
		}
		sort.Ints(arr)
		for i := index; i < n; i++ {
			fmt.Print(arr[i], " ")
		}
		for i := 0; i < len(last); i++ {
			fmt.Print(last[i], " ")
		}
		fmt.Println()
	}

}

func NewCF1497A(r *bufio.Reader) *CF1497A {
	return &CF1497A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1497A(bufio.NewReader(os.Stdin)).Run()
}
