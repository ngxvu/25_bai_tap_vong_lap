// Nhập số nguyên dương n bất kỳ. Viết chương trình vẽ tam giác cân "*" rỗng có chiều cao là n hàng.

package main

import (
	"fmt"
)

func main() {
	var n int = 4
	if n > 0 {
		for i := 0; i < n; i++ {
			for j := 0; j < 2*n+1; j++ {
				if j == n-i || j == n+i {
					fmt.Print(" * ")
				} else {
					fmt.Print("   ")
				}
			}
			fmt.Print(("\n"))
		}
		for j := 0; j < 2*n+1; j++ {
			fmt.Print(" * ")
		}

	}
}
