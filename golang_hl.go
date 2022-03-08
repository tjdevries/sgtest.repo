package main

import "fmt"

func main() {}

func highlightContent() {
	x := "hello"

	y := func(x string) {
		fmt.Println(x)
	}

	y(x)
}
