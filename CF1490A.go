package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1490A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1490A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1490A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1490A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1490A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1490A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1490A
 Date: 25/05/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1490/A
 Problem: CF1490A
**/
func (in *CF1490A) Run() {

	t := in.NextInt()

	for ; t > 0; t-- {
		//n := in.NextInt()

	}

}

func NewCF1490A(r *bufio.Reader) *CF1490A {
	return &CF1490A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1490A(bufio.NewReader(os.Stdin)).Run()
}
