// Nhập 2 số nguyên x và y. Viết chương trình tính tổng bình phương các số từ x đến y.

package main

import "fmt"

func main() {
	var x = 3
	var y = 7
	var S int
	for i := x; i <= y; i++ {
		S = S + i*i
	}
	fmt.Print(S)
}
