package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1534C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1534C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1534C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1534C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1534C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1534C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1534C
 Date: 13/06/21
 User: pepradere
 URL: https://codeforces.com/contest/1534/problem/C
 Problem: CF1534C
**/
func (in *CF1534C) Run() {

	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		a := []int{}
		b := []int{}

		for i := 0; i < n; i++ {
			a = append(a, in.NextInt()-1)
		}

		for i := 0; i < n; i++ {
			b = append(b, in.NextInt()-1)
		}

		mady := []map[int]bool{}
		for i := 0; i < n; i++ {
			mady = append(mady, map[int]bool{})
		}

		for i := 0; i < n; i++ {
			mady[a[i]][b[i]] = true
			mady[b[i]][a[i]] = true
		}

		cola := []int{}
		vis := []bool{}
		for i := 0; i <= n; i++ {
			cola = append(cola, 0)
			vis = append(vis, false)
		}

		count := 0
		for i := 0; i < n; i++ {
			if !vis[i] {
				u := i
				vis[u] = true
				p, c := 0, 0
				cola[c] = u
				c++
				for p < c {
					u = cola[p]
					p++
					for v, _ := range mady[u] {
						if !vis[v] {
							vis[v] = true
							cola[c] = v
							c++
						}

					}
				}
				count++
			}
		}

		const S = 1000000007
		ans := 1
		for i := 0; i < count; i++ {
			ans = ((ans % S) * 2) % S
		}

		//ans := int64(math.Pow(2.0, float64(count))) % 1000000007
		fmt.Println(ans)

	}
}

func NewCF1534C(r *bufio.Reader) *CF1534C {
	return &CF1534C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1534C(bufio.NewReader(os.Stdin)).Run()
}
