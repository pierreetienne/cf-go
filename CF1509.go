package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1509 struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1509) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1509) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1509) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1509) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1509) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1509
 Date: 17/04/21
 User: pepradere
 URL: https://codeforces.com/contest/1509/problem/A
 Problem: CF1509
**/
func (in *CF1509) Run() {
	T := in.NextInt()
	for ; T > 0; T-- {
		N := in.NextInt()
		var even []int
		var odd []int
		for ; N > 0; N-- {
			num := in.NextInt()
			if num%2 == 0 {
				even = append(even, num)
			} else {
				odd = append(odd, num)
			}
		}
		var sol []int
		if len(even) > len(odd) {
			sol = append(sol, even...)
			sol = append(sol, odd...)
		} else {
			sol = append(sol, odd...)
			sol = append(sol, even...)
		}

		for i := 0; i < len(sol); i++ {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(sol[i])
		}
		fmt.Println()
	}

}

func NewCF1509(r *bufio.Reader) *CF1509 {
	return &CF1509{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1509(bufio.NewReader(os.Stdin)).Run()
}
