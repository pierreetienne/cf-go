package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1473B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1473B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1473B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1473B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1473B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1473B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1473B
 Date: 11/06/21
 User: pepradere
 URL: https://codeforces.com/contest/1473/problem/B
 Problem: CF1473B
**/
func (in *CF1473B) Run() {

	q := in.NextInt()
	for ; q > 0; q-- {
		s := in.NextString()
		t := in.NextString()
		ans := "-1"
		if len(s) == len(t) {
			if s == t {
				ans = s
			}
		} else {
			ans = in.g(s, t)
		}

		fmt.Println(ans)
	}

}

func (in *CF1473B) g(s, t string) string {
	os := s
	ot := t
	ans := "-1"
	a, b := 1, 1
	for len(s) != len(t) && a < 5000 && b < 5000 {
		if len(s) < len(t) {
			s = in.f(s, os)
			a++
		} else if len(s) > len(t) {
			t = in.f(t, ot)
			b++
		}
	}
	if len(s) == len(t) && s == t {
		ans = t
	}

	return ans
}

func (in *CF1473B) f(str string, o string) string {
	return o + str
}

func NewCF1473B(r *bufio.Reader) *CF1473B {
	return &CF1473B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1473B(bufio.NewReader(os.Stdin)).Run()
}
