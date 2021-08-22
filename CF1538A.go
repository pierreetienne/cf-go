package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1538A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1538A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1538A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1538A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1538A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1538A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1538A
 Date: 18/08/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1538/A
 Problem: CF1538A
**/
func (in *CF1538A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		min, max := n+1, 0
		posMin, posMax := -1, -1
		for i := 1; i <= n; i++ {
			num := in.NextInt()
			if num < min {
				min = num
				posMin = i
			}
			if num > max {
				max = num
				posMax = i
			}
		}

		ans := 0

		if posMax < posMin {
			posMax, posMin = posMin, posMax
		}

		n++
		if posMin < n-posMax {
			ans += posMin
			if posMax-posMin < n-posMax {
				ans += posMax - posMin
			} else {
				ans += n - posMax
			}
		} else {
			ans += n - posMax
			if posMax-posMin < posMin {
				ans += posMax - posMin
			} else {
				ans += posMin
			}
		}
		fmt.Println(ans)
	}
}

func NewCF1538A(r *bufio.Reader) *CF1538A {
	return &CF1538A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1538A(bufio.NewReader(os.Stdin)).Run()
}
