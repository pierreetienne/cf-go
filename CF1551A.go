package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1551A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1551A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1551A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1551A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1551A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1551A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1551A
 Date: 18/08/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1551/A
 Problem: CF1551A
**/
func (in *CF1551A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		two := int(math.Ceil(float64(n) / 3.0))
		one := int(math.Floor(float64(n) / 3.0))

		if (two * 2) + one != n {
			one, two = two, one
		}
		fmt.Println(one, two)
	}
}

func NewCF1551A(r *bufio.Reader) *CF1551A {
	return &CF1551A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1551A(bufio.NewReader(os.Stdin)).Run()
}
