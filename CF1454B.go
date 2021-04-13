package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1454B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1454B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1454B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1454B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1454B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1454B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1454B
 Date: 12/04/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1454/B
 Problem: CF1454B
**/
func (in *CF1454B) Run() {
	t := in.NextInt()
	//fmt.Println("t : " , t)
	for ; t > 0; t-- {
		n := in.NextInt()
		//fmt.Println("n : " , n)
		count := make(map[int]int)
		time := make(map[int]int)
		for i := 0; n > 0; n-- {
			value := in.NextInt()
			//fmt.Println("value : " , value)
			c, ok := count[value]
			if !ok {
				c = 0
				count[value] = 0
				time[value] = i
			}
			count[value] = c + 1
			i++
		}
		minTime, ans := 10000000, -1
		for key := range count {
			c := count[key]
			if c == 1 {
				tc := time[key]
				if ans == -1 || ans > key {
					ans = key
					minTime = tc
				}
			}
		}
		if ans == -1 {
			fmt.Println(-1)
		} else {
			fmt.Println(minTime + 1)
		}
	}

}

func NewCF1454B(r *bufio.Reader) *CF1454B {
	return &CF1454B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1454B(bufio.NewReader(os.Stdin)).Run()
}
