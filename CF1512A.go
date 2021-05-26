package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1512A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1512A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1512A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1512A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1512A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1512A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1512A
 Date: 16/05/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1512/A
 Problem: CF1512A
**/
func (in *CF1512A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		data := map[int][]int{}
		for i := 0; n > 0; n-- {
			val := in.NextInt()
			_, ok := data[val]
			if !ok {
				data[val] = []int{0, i}
			}
			data[val][0] = data[val][0] + 1
			i++
		}

		for d := range data {
			if data[d][0] == 1 {
				fmt.Println(data[d][1] + 1)
			}
		}

	}
}

func NewCF1512A(r *bufio.Reader) *CF1512A {
	return &CF1512A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1512A(bufio.NewReader(os.Stdin)).Run()
}
