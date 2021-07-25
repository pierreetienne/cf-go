package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1360C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1360C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1360C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1360C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1360C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1360C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1360C
 Date: 24/07/21
 User: pepradere
 URL: https://codeforces.co1m/problemset/problem/1360/C
 Problem: CF1360C
**/
func (in *CF1360C) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)

		even := 0
		for i := 0; i < n ; i++ {
			arr[i] = in.NextInt()
			if arr[i]%2 == 0 {
				even++
			}
		}
 		odd := n - even

		ws := false

		sort.Ints(arr)
		if even % 2 == 0 && odd % 2 == 0 {
			ws = true
		} else if even > 0 && odd > 0 {
			consecutive := false
			for i:=0;i<n-1;i++ {
				if arr[i]+1 == arr[i+1] {
					consecutive = true
					break
				}
			}
			if consecutive {
				ws = true
			}
		}

		if ws {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}


	}

}

func NewCF1360C(r *bufio.Reader) *CF1360C {
	return &CF1360C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1360C(bufio.NewReader(os.Stdin)).Run()
}
