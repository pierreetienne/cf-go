package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1388A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1388A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1388A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1388A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1388A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1388A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1388A
 Date: 2/04/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1388/A
 Problem: CF1388A
**/
func (in *CF1388A) Run() {

	const P = 10

	var primes [P]bool
	for i := 2; i < P; i++ {
		for j := i * 2; j < P; j += i {
			primes[j] = true
		}
	}

	var primesValues []int
	for i := 2; i < P; i++ {
		if !primes[i] {
			primesValues = append(primesValues, i)
		}
	}

	nearlyPrimes := make(map[int]bool)

	for i := 0; i < len(primesValues); i++ {
		for j := i + 1; j < len(primesValues); j++ {
			val := primesValues[i] * primesValues[j]
			if val <= 200000 {
				nearlyPrimes[val] = true
			} else {
				break
			}
		}
	}

	var values []int

	for near := range nearlyPrimes {
		values = append(values, near)
	}

	sort.Ints(values)

	T := in.NextInt()
	for ; T > 0; T-- {
		N := in.NextInt()
		var ws = false
		for _, i := range values {
			if i > N || ws {
				break
			}
			for _, j := range values {
				if i+j > N || ws {
					break
				}
				for _, p := range values {
					if i+j+p > N || ws {
						break
					}
					val := N - i - j - p
					if i < j && j < p && i+j+p < N && val != i && val != j && val != p {
						ws = true
						fmt.Println("YES")
						fmt.Println(i, j, p, val)
					}
				}
			}
		}
		if !ws {
			fmt.Println("NO")
		}

	}

}

func NewCF1388A(r *bufio.Reader) *CF1388A {
	return &CF1388A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1388A(bufio.NewReader(os.Stdin)).Run()
}
