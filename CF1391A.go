package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1391A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1391A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1391A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1391A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1391A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1391A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1391A
 Date: 3/04/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1391/A
 Problem: CF1391A
**/
func (in *CF1391A) Run() {
	T := in.NextInt()

	for ; T > 0; T-- {
		N := in.NextInt()
		for i := 0; i < N; i++ {
			if i == 0 {
				fmt.Print(i + 1)
			} else {
				fmt.Print(" ", i+1)
			}
		}
		fmt.Println()
	}

}

func NewCF1391A(r *bufio.Reader) *CF1391A {
	return &CF1391A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1391A(bufio.NewReader(os.Stdin)).Run()
}
