package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1534B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1534B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1534B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1534B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1534B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1534B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 WA
 Run solve the problem CF1534B
 Date: 13/06/21
 User: pepradere
 URL: https://codeforces.com/contest/1534/problem/B
 Problem: CF1534B
**/
func (in *CF1534B) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		var arr []int
		for i := 0; i < n; i++ {
			arr = append(arr, in.NextInt())
		}

		aux := 0
		sol := calc(arr)
		for {
			init := calc(arr) + aux
			cpy := []int{}
			for i := 0; i < len(arr); i++ {
				cpy = append(cpy, arr[i])
			}
			for i := 0; i < n; i++ {
				if arr[i] > 0 {
					if i == 0 {
						if i+1 < n {
							if arr[i] > arr[i+1] {
								cpy[i]--
								aux++
							}
						}
					} else if i == n-1 {
						if i-1 >= 0 {
							if arr[i] > arr[i-1] {
								cpy[i]--
								aux++
							}
						}
					} else {
						if arr[i] >= arr[i+1] && arr[i] >= arr[i-1] {
							cpy[i]--
							aux++
						}
					}
				}
			}

			arr = cpy
			fin := calc(arr) + aux

			if fin < sol {
				sol = fin
			}
			if fin == init {

				break
			}
		}
		fmt.Println(sol)

	}

}

func calc(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			sum += arr[i]
		} else {
			if arr[i-1] > arr[i] {
				sum += arr[i-1] - arr[i]
			} else {
				sum += arr[i] - arr[i-1]
			}
		}
	}
	return sum + arr[len(arr)-1]
}

func NewCF1534B(r *bufio.Reader) *CF1534B {
	return &CF1534B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1534B(bufio.NewReader(os.Stdin)).Run()
}
