package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cFRead struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (read *cFRead) GetLine() string {
	line, err := read.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	read.split = []string{}
	read.index = 0
	return line
}
func (read *cFRead) load() {
	if len(read.split) <= read.index {
		read.split = strings.Split(read.GetLine(), read.separator)
		read.index = 0
	}
}

func (read *cFRead) NextInt() int {
	read.load()
	val, _ := strconv.Atoi(strings.TrimSpace(read.split[read.index]))
	read.index++
	return val
}

func (read *cFRead) NextInt64() int64 {
	read.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(read.split[read.index]), 10, 64)
	read.index++
	return val
}

func (read *cFRead) NextString() string {
	read.load()
	val := strings.TrimSpace(read.split[read.index])
	read.index++
	return val
}

func newCFRead(r *bufio.Reader) *CFRead {
	return &CFRead{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	in := newCFRead(bufio.NewReader(os.Stdin))

	in.NextInt()

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	line := text[:len(text)-1]

	fmt.Println("last ", line, " last ", line[len(line)-1:])
	var t int = 0

	fmt.Scan(&t)
	fmt.Println("T ", t)
	for ; t > 0; t-- {
		var n int = 0
		fmt.Scan(&n)
		fmt.Println("N ", n)
		for ; n > 0; n-- {
			var x int = 0
			fmt.Scan(&x)
			fmt.Println("val ", x)
		}
	}

}
