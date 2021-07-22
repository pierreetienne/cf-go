package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF964A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF964A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF964A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF964A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF964A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF964A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF964A
 Date: 21/07/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/964/A
 Problem: CF964A
**/
func (in *CF964A) Run() {
	N := in.NextInt()
	sol := 1
	if N > 1 {
		sol = (N / 2) + 1
	}
	fmt.Println(sol)
}

func NewCF964A(r *bufio.Reader) *CF964A {
	return &CF964A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF964A(bufio.NewReader(os.Stdin)).Run()
}
