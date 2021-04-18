package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1509C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
	arr       []int
	memo      [][]int64
}

func (in *CF1509C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1509C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1509C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1509C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1509C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1509C
 Date: 17/04/21
 User: pepradere
 URL: https://codeforces.com/contest/1509/problem/C
 Problem: CF1509C
**/
func (in *CF1509C) Run() {
	N := in.NextInt()

	for ; N > 0; N-- {
		in.arr = append(in.arr, in.NextInt())
	}
	sort.Ints(in.arr)
	in.memo = make([][]int64, len(in.arr))
	for i := 0; i < len(in.arr); i++ {
		in.memo[i] = make([]int64, len(in.arr))
		for j := 0; j < len(in.arr); j++ {
			in.memo[i][j] = -1
		}
	}
	sol := in.dp(0, len(in.arr)-1)
	fmt.Println(sol)

}

func (in *CF1509C) dp(f, t int) int64 {
	if f == t {
		return 0
	}
	value := int64(in.arr[t] - in.arr[f])
	if in.memo[f][t] == -1 {
		in.memo[f][t] = value + int64(math.Min(float64(in.dp(f+1, t)), float64(in.dp(f, t-1))))
	}
	return in.memo[f][t]
}

func NewCF1509C(r *bufio.Reader) *CF1509C {
	return &CF1509C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1509C(bufio.NewReader(os.Stdin)).Run()
}
