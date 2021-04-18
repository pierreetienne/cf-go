package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1509B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1509B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1509B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1509B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1509B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1509B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1509B
 Date: 17/04/21
 User: pepradere
 URL: https://codeforces.com/contest/1509/problem/B
 Problem: CF1509B
**/
func (in *CF1509B) Run() {
	T := in.NextInt()
	for ; T > 0; T-- {
		in.NextInt()
		cad := in.NextString()
		M, P := 0, 0
		ws := true
		for i := 0; i < len(cad); i++ {
			if cad[i] == 'M' {
				M++
			} else {
				P++
			}
			if M > P {
				ws = false
			}
		}
		M = 0
		P = 0
		for i := len(cad) - 1; i >= 0; i-- {
			if cad[i] == 'M' {
				M++
			} else {
				P++
			}
			if M > P {
				ws = false
			}
		}

		if M*2 != P || cad[len(cad)-1] == 'M' {
			ws = false
		}

		if ws {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1509B(r *bufio.Reader) *CF1509B {
	return &CF1509B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1509B(bufio.NewReader(os.Stdin)).Run()
}
