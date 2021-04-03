package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1473A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1473A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1473A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1473A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1473A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1473A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1473A
 Date: 2/04/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1473/A
 Problem: CF1473A
**/
func (in *CF1473A) Run() {
	T := in.NextInt()
	for ; T > 0; T-- {
		N := in.NextInt()
		D := in.NextInt()
		var arr [100]int
		var ws bool = false
		i := 0
		for ; N > 0; N-- {
			value := in.NextInt()
			if value > D {
				ws = true
			}
			arr[i] = value
			i++
		}

		if !ws {
			fmt.Println("YES")
		} else {
			sort.Ints(arr[:i])
			if arr[0]+arr[1] <= D {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}

}

func NewCF1473A(r *bufio.Reader) *CF1473A {
	return &CF1473A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1473A(bufio.NewReader(os.Stdin)).Run()
}
