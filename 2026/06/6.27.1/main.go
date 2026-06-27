package main

import "fmt"

func main() {
	const (
		err1 = iota + 4
		err2
		err3 = "ha"
		err25
		err4
		err5 = iota
		err6
	)
	fmt.Println(err1, err2, err3, err25, err4, err5, err6)

	const (
		err11 = iota
		err12
		err13
		err14
	)
	fmt.Println(err11, err12, err13, err14)
}
