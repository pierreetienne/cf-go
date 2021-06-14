package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1534A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1534A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1534A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1534A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1534A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1534A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1534A
 Date: 13/06/21
 User: pepradere
 URL: https://codeforces.com/contest/1534/problem/0
 Problem: CF1534A
**/
func (in *CF1534A) Run() {
	t := in.NextInt()

	for ; t > 0; t-- {
		n := in.NextInt()
		m := in.NextInt()
		var arr []string
		for i := 0; i < n; i++ {
			arr = append(arr, in.NextString())
		}

		indexI := -1
		indexJ := -1
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if arr[i][j] != '.' {
					indexI = i
					indexJ = j
					break
				}
			}
			if indexI != -1 {
				break
			}
		}

		ans := "YES"
		var vis []bool
		for i := 0; i < n; i++ {
			vis = append(vis, false)
		}

		if indexI == -1 && indexJ == -1 {
			indexI = 0
			indexJ = 0
			arr[0] = replaceAtIndex2(arr[0], 'R', 0)
		}

		last := arr[indexI][indexJ]
		for j := indexJ + 1; j < m; j++ {
			last = other(last)
			if arr[indexI][j] == '.' {
				arr[indexI] = replaceAtIndex2(arr[indexI], last, j)
			} else if arr[indexI][j] != last {
				ans = "NO"
				break
			} else {
				continue
			}
		}

		last = arr[indexI][indexJ]
		for j := indexJ - 1; j >= 0; j-- {
			last = other(last)
			if arr[indexI][j] == '.' {
				arr[indexI] = replaceAtIndex2(arr[indexI], last, j)
			} else if arr[indexI][j] != last {
				ans = "NO"
				break
			} else {
				continue
			}
		}

		for i := indexI - 1; i >= 0; i-- {
			last = arr[i+1][0]
			for j := 0; j < m; j++ {
				last = other(last)
				if arr[i][j] == '.' {
					arr[i] = replaceAtIndex2(arr[i], last, j)
				} else if arr[i][j] != last {
					ans = "NO"
					break
				} else {
					continue
				}
			}
			if ans == "NO" {
				break
			}
		}

		for i := indexI + 1; i < n; i++ {
			last = arr[i-1][0]
			for j := 0; j < m; j++ {
				last = other(last)
				if arr[i][j] == '.' {
					arr[i] = replaceAtIndex2(arr[i], last, j)
				} else if arr[i][j] != last {
					ans = "NO"
					break
				} else {
					continue
				}
			}
			if ans == "NO" {
				break
			}
		}

		fmt.Println(ans)
		if ans == "YES" {
			for i := 0; i < n; i++ {
				fmt.Println(arr[i])
			}
		}

	}

}

func other(x uint8) uint8 {
	if x == 'R' {
		return 'W'
	}
	return 'R'
}

func replaceAtIndex2(str string, replacement uint8, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func NewCF1534A(r *bufio.Reader) *CF1534A {
	return &CF1534A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1534A(bufio.NewReader(os.Stdin)).Run()
}
