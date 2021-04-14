package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF978A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF978A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF978A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF978A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF978A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF978A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF978A
 Date: 13/04/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/978/A
 Problem: CF978A
**/
func (in *CF978A) Run() {
	N := in.NextInt()
	var arr []int
	for ; N > 0; N-- {
		arr = append(arr, in.NextInt())
	}
	m := make(map[int]bool)
	var res []int
	for i := len(arr) - 1; i >= 0; i-- {
		_, ok := m[arr[i]]
		if !ok {
			m[arr[i]] = true
			res = append(res, arr[i])
		}
	}
	fmt.Println(len(res))
	for i := len(res) - 1; i >= 0; i-- {
		if i == len(res)-1 {
			fmt.Print(res[i])
		} else {
			fmt.Print(" ", res[i])
		}

	}
}

func NewCF978A(r *bufio.Reader) *CF978A {
	return &CF978A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF978A(bufio.NewReader(os.Stdin)).Run()
}
