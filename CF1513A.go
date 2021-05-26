package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1513A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1513A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1513A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1513A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1513A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1513A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1513A
 Date: 21/04/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1513/A
 Problem: CF1513A
**/
func (in *CF1513A) Run() {
	N := in.NextInt()
	for ; N > 0; N-- {
		n := in.NextInt()
		k := in.NextInt()
		var arr []int
		for i := 0; i < n; i++ {
			arr = append(arr, i+1)
		}
		count := 0
		for i := 2; i < n && count < k; i += 2 {
			aux := arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = aux
			count++
		}

		if count == k {
			for i := 0; i < n; i++ {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(arr[i])
			}
			fmt.Println()
		} else {
			fmt.Println(-1)
		}
	}
}

func NewCF1513A(r *bufio.Reader) *CF1513A {
	return &CF1513A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1513A(bufio.NewReader(os.Stdin)).Run()
}
