package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

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
