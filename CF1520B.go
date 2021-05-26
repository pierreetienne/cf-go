package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1520B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1520B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1520B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1520B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1520B) NextInt64() uint64 {
	in.load()
	val, _ := strconv.ParseUint(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1520B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
 Run solve the problem CF1520B
 Date: 24/05/21
 User: pepradere
 URL: https://codeforces.com/problemset/problem/1520/B
 Problem: CF1520B
**/
func (in *CF1520B) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextString()
		if len(n) == 1 {
			fmt.Println(n)
		} else {
			input, _ := strconv.Atoi(n)
			ans := 0
			for i := 1; i <= input; i = i*10 + 1 {
				for j := 1; j <= 9; j++ {
					if j*i <= input {
						ans++
					}
				}
			}

			fmt.Println(ans)
			//for j := 0; j < len(n); j++ {
			//	num := uint8(9)
			//	if j == 0 {
			//		num = n[j] - '0'
			//	}
			//	ws := false
			//	wsStr := ""
			//	for ; num > 0; num-- {
			//		str := string(num + '0')
			//		for i := 0; i < len(n)-j-1; i++ {
			//			str += string(num + '0')
			//		}
			//		strNum, _ := strconv.ParseInt(str, 10, 64)
			//		if strNum <= input {
			//			ws = true
			//			wsStr = str
			//			break
			//		}
			//	}
			//	if ws {
			//		sum := 9
			//		for i := 0; i < len(wsStr)-1; i++ {
			//			if i == 0 {
			//				sum += int(wsStr[i] - '0')
			//			} else {
			//				sum += 9
			//			}
			//		}
			//		fmt.Println(sum)
			//		break
			//	}
			//}
		}
	}

}

func NewCF1520B(r *bufio.Reader) *CF1520B {
	return &CF1520B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1520B(bufio.NewReader(os.Stdin)).Run()
}
