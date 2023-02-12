// Viết chương trình in bảng cửu chương rút gọn của các số từ 1 đến 100.

package main

import "fmt"

func main() {
	fmt.Println("In bảng cửu chương rút gọn của các số từ 1 đến 100.")
	for x := 1; x <= 100; x++ {
		for y := 1; y <= 10; y++ {
			fmt.Print(x*y, "  ")
		}
		fmt.Print("\n")
	}
}
