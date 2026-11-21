package main

import "fmt"

// func main() {
// 	var x int
// 	x = 5
// 	fmt.Println(x * 2)
// }

func main() {
	score := 20

	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
