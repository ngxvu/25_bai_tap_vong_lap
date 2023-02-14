// Viết chương trình nhập vào 1 số nguyên n và in tất cả các số từ n đến -100.
// Bắt nhập lại nếu người dùng nhập -1.

package main

import "fmt"

func main() {
	var n int = -1
	if n > -100 {
		for i := n; i >= -100; i-- {
			fmt.Print(i, "\n")
		}
	} else {
		for i := n; i <= -100; i++ {
			fmt.Print(i, "\n")
		}
	}
}
