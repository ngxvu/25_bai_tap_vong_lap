// Viết chương trình in bảng số từ 1 đến 200.

package main

import "fmt"

func main() {
	fmt.Println("In bảng số từ 1 đến 200.")
	for x := 1; x <= 10; x++ {
		for y := x; y <= 200; y += 10 {
			fmt.Print(y, " ")
		}
		fmt.Print("\n")
	}
}
