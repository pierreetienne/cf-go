package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CFRead struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (read *CFRead) GetLine() string {
	line, err := read.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	read.split = []string{}
	read.index = 0
	return line
}
func (read *CFRead) load() {
	if len(read.split) <= read.index {
		read.split = strings.Split(read.GetLine(), read.separator)
		read.index = 0
	}
}

func (read *CFRead) NextInt() int {
	read.load()
	val, _ := strconv.Atoi(strings.TrimSpace(read.split[read.index]))
	read.index++
	return val
}

func (read *CFRead) NextInt64() int64 {
	read.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(read.split[read.index]), 10, 64)
	read.index++
	return val
}

func (read *CFRead) NextString() string {
	read.load()
	val := strings.TrimSpace(read.split[read.index])
	read.index++
	return val
}

func NewCFRead(r *bufio.Reader) *CFRead {
	return &CFRead{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	in := NewCFRead(bufio.NewReader(os.Stdin))

	T := in.NextInt()
	for ; T > 0; T-- {
		N := in.NextInt()
		var arr []int
		for ; N > 0; N-- {
			arr = append(arr, in.NextInt())
		}
		if arr[0]+arr[1] <= arr[len(arr)-1] {
			fmt.Println("1 2", len(arr))
		} else {
			fmt.Println("-1")
		}

	}
}
