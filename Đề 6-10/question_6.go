// Nhập số nguyên n bất kỳ. Viết chương trình in các số chẵn từ 1 đến n.

package main

import (
	"fmt"
)

func main() {
	var x int = 6
	for i := 1; i <= x; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}
