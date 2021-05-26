package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1520A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1520A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1520A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1520A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1520A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1520A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1520A
 Date: 16/05/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1520/A
 Problem: CF1520A
**/
func (in *CF1520A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		line := in.NextString()
		data := map[uint8]int{}
		ws := false
		for i := 0; n > 0; n-- {
			index, ok := data[line[i]]
			if ok && index+1 != i {
				fmt.Println("NO")
				ws = true
				break
			}
			data[line[i]] = i
			i++
		}
		if !ws {
			fmt.Println("YES")
		}
	}
}

func NewCF1520A(r *bufio.Reader) *CF1520A {
	return &CF1520A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1520A(bufio.NewReader(os.Stdin)).Run()
}
