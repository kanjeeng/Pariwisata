package main

import "fmt"

func main() {
	for i := 1; i < 10000; i++ {
		if (45*i)-135%947 == 181 {
			fmt.Print(i)
		}
	}
}
