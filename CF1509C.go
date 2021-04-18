package main

import (
	"bufio"
	"fmt"
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
	// TODO WA
	N := in.NextInt()
	var arr []int
	for ; N > 0; N-- {
		arr = append(arr, in.NextInt())
	}
	sort.Ints(arr)

	mov := 0
	for i := 0; i < len(arr); i++ {
		if arr[0] == arr[i] {
			arr = append(arr, arr[0])
			mov++
		} else {
			break
		}
	}
	if mov > 0 {
		arr = arr[mov:]
	}
	//fmt.Println("---- " , arr)
	max := 0
	min := 2000000000
	d := 0
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
		d += max - min
	}
	fmt.Println(d)
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
