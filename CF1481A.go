package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1481A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1481A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1481A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1481A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1481A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1481A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1481A
 Date: 26/05/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1481/A
 Problem: CF1481A
**/
func (in *CF1481A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		px := in.NextInt()
		py := in.NextInt()
		chain := in.NextString()

		data := map[uint8]int{}
		ws1, ws2 := false, false
		for i := 0; i < len(chain); i++ {
			val, ok := data[chain[i]]
			if !ok {
				data[chain[i]] = 0
			}
			data[chain[i]] = val + 1
		}

		U, _ := data['U']
		D, _ := data['D']
		R, _ := data['R']
		L, _ := data['L']

		if px < 0 && L >= -px {
			ws1 = true
		}
		if px >= 0 && R >= px {
			ws1 = true
		}

		if py < 0 && D >= -py {
			ws2 = true
		}
		if py >= 0 && U >= py {
			ws2 = true
		}

		if ws1 && ws2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}

}

func NewCF1481A(r *bufio.Reader) *CF1481A {
	return &CF1481A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1481A(bufio.NewReader(os.Stdin)).Run()
}
