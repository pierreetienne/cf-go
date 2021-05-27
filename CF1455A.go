package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1455A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1455A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1455A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1455A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1455A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1455A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1455A
 Date: 26/05/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1455/A
 Problem: CF1455A
**/
func (in *CF1455A) Run() {

	t := in.NextInt()
	for ; t > 0; t-- {
		fmt.Println(len(in.NextString()))
	}
}

func NewCF1455A(r *bufio.Reader) *CF1455A {
	return &CF1455A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1455A(bufio.NewReader(os.Stdin)).Run()
}
