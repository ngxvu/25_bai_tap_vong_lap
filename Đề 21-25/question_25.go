// Nhập số nguyên n. Tính giá trị biểu thức S= 1.2 + 2.3 + 3.4 + ... + n(n+1).

package main

import "fmt"

func main() {
	var n int = 3
	var s int = 0
	for i := 1; i <= n; i++ {
		s = s + i*(i+1)
	}
	fmt.Print(s)
}
