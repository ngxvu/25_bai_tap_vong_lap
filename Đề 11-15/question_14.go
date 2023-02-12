// Nhập số nguyên dương x bất kỳ. Viết chương trình tính giai thừa của x.

package main

import "fmt"

func main() {
	var n int = 5
	var gt int = 1
	for i := 1; i <= n; i++ {
		gt = gt * i

	}
	fmt.Print("giai thua cua ", n, " la ", gt)
}
