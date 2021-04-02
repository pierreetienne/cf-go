package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type cf1473aRead struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (read *cf1473aRead) GetLine() string {
	line, err := read.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	read.split = []string{}
	read.index = 0
	return line
}
func (read *cf1473aRead) load() {
	if len(read.split) <= read.index {
		read.split = strings.Split(read.GetLine(), read.separator)
		read.index = 0
	}
}

func (read *cf1473aRead) NextInt() int {
	read.load()
	val, _ := strconv.Atoi(strings.TrimSpace(read.split[read.index]))
	read.index++
	return val
}

func (read *cf1473aRead) NextInt64() int64 {
	read.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(read.split[read.index]), 10, 64)
	read.index++
	return val
}

func (read *cf1473aRead) NextString() string {
	read.load()
	val := strings.TrimSpace(read.split[read.index])
	read.index++
	return val
}

func newcf1473aread(r *bufio.Reader) *cf1473aRead {
	return &cf1473aRead{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	in := newcf1473aread(bufio.NewReader(os.Stdin))

	T := in.NextInt()
	for ; T > 0; T-- {
		N := in.NextInt()
		D := in.NextInt()
		var arr [100]int
		var ws bool = false
		i := 0
		for ; N > 0; N-- {
			value := in.NextInt()
			if value > D {
				ws = true
			}
			arr[i] = value
			i++
		}

		if !ws {
			fmt.Println("YES")
		} else {
			sort.Ints(arr[:i])
			if arr[0]+arr[1] <= D {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}

	}
}
