package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1398B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1398B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1398B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1398B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1398B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1398B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1398B
 Date: 13/04/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1398/B
 Problem: CF1398B
**/
func (in *CF1398B) Run() {

	T := in.NextInt()
	for ; T > 0; T-- {
		cad := in.NextString()
		var cont = -1
		var arr []int
		for i := 0; i < len(cad); i++ {
			if cad[i] == '1' {
				if cont == -1 {
					cont = 0
				}
				cont++
			} else if cont != -1 {
				arr = append(arr, cont)
				cont = -1
			}
		}
		if cont != -1 {
			arr = append(arr, cont)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(arr)))

		alice := 0

		for i := 0; i < len(arr); i++ {
			if i%2 == 0 {
				alice += arr[i]
			}
		}

		fmt.Println(alice)

	}
}

func NewCF1398B(r *bufio.Reader) *CF1398B {
	return &CF1398B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1398B(bufio.NewReader(os.Stdin)).Run()
}
