package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1535A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1535A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1535A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1535A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1535A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1535A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1535A
 Date: 11/06/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1535/A
 Problem: CF1535A
**/
func (in *CF1535A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		s1, s2, s3, s4 := in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()

		arr := []int{s1, s2, s3, s4}
		sort.Ints(arr)

		ans := "YES"
		if (arr[3] == s1 && arr[2] == s2) || (arr[3] == s2 && arr[2] == s1) {
			ans = "NO"
		} else if (arr[3] == s3 && arr[2] == s4) || (arr[3] == s4 && arr[2] == s3) {
			ans = "NO"
		}

		fmt.Println(ans)
	}
}

func NewCF1535A(r *bufio.Reader) *CF1535A {
	return &CF1535A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1535A(bufio.NewReader(os.Stdin)).Run()
}
