package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1542A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1542A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1542A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1542A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1542A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1542A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/** WA*******
 Run solve the problem CF1542A
 Date: 21/08/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1542/A
 Problem: CF1542A
**/
func (in *CF1542A) Run() {
	n := in.NextInt()
	for ; n > 0; n-- {
		t := in.NextInt()
		odd := 0
		even := 0
		for i := 0; i < 2*t; i++ {
			num := in.NextInt()
			if num %2 == 0{
				even++
			} else {
				odd++
			}
		}
		if odd == even && odd == t {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func NewCF1542A(r *bufio.Reader) *CF1542A {
	return &CF1542A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1542A(bufio.NewReader(os.Stdin)).Run()
}
