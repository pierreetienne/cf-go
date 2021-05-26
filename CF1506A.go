package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1506A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1506A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1506A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1506A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1506A) NextInt64() uint64 {
	in.load()
	val, _ := strconv.ParseUint(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1506A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1506A
 Date: 16/05/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1506/A
 Problem: CF1506A
**/
func (in *CF1506A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n, m, x := in.NextInt64(), in.NextInt64(), in.NextInt64()
		x--
		a := x / n
		b := x % n
		res := (b * m) + a
		fmt.Println(res + 1)
	}
}

func NewCF1506A(r *bufio.Reader) *CF1506A {
	return &CF1506A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1506A(bufio.NewReader(os.Stdin)).Run()
}
