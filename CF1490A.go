package main

import (
	"bufio"
	"fmt"
	"math"
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

func (in *CF1490A) Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (in *CF1490A) Max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
		n := in.NextInt()
		var arr []int
		for ; n > 0; n-- {
			val := in.NextInt()
			arr = append(arr, val)
		}

		var sol []int
		for i := 0; i < len(arr); i++ {
			if i > 0 {
				for in.Min(sol[len(sol)-1], arr[i])*2 < in.Max(sol[len(sol)-1], arr[i]) {
					if arr[i] == in.Min(sol[len(sol)-1], arr[i]) {
						sol = append(sol, int(math.Ceil(float64(sol[len(sol)-1])/2.0)))
					} else {
						sol = append(sol, sol[len(sol)-1]*2)
					}
				}
			}
			sol = append(sol, arr[i])
		}
		fmt.Println(len(sol) - len(arr))
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
