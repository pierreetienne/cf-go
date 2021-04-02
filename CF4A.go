package main

import "fmt"

func main() {
	var val int = 0
	fmt.Scan(&val)
	if val%2 == 0 && val != 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
