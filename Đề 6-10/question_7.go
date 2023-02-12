// Nhập số nguyên n bất kỳ. Viết chương trình in các số lẻ từ 1 đến n.

package main

import "fmt"

func main() {
	var x int = 10
	fmt.Println("Các số chẵn từ 1 đến", x, "bao gom: ")
	for i := 1; i <= x; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}
}
