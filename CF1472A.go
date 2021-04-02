package main

import (
	"fmt"
)

func main() {

	T := 0
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var w, h, n = 0, 0, 0
		fmt.Scan(&w, &h, &n)

		var total = 1
		for w%2 == 0 {
			w /= 2
			total *= 2
		}

		for h%2 == 0 {
			h /= 2
			total *= 2
		}

		if total >= n {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
