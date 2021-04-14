package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1478A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1478A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1478A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1478A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1478A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1478A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1478A
 Date: 13/04/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1478/A
 Problem: CF1478A
**/
func (in *CF1478A) Run() {
	T := in.NextInt()
	for ; T > 0; T-- {
		N := in.NextInt()
		m := make(map[int]int)
		max := 0
		for ; N > 0; N-- {
			val := in.NextInt()
			d, ok := m[val]
			if !ok {
				m[val] = 1
			} else {
				m[val] = d + 1
			}
			d = m[val]
			max = int(math.Max(float64(d), float64(max)))
		}

		fmt.Println(max)
	}
}

func NewCF1478A(r *bufio.Reader) *CF1478A {
	return &CF1478A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1478A(bufio.NewReader(os.Stdin)).Run()
}
