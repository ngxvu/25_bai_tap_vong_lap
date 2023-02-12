// Nhập số nguyên n bất kỳ, viết chương trình in các phần tử đầu tiên của dãy Fibonacci từ 0 đến n.
// input 20: 0 1 1 2 3 5 8 13

package main

import "fmt"

func main() {
	Solution(20)
}

func Solution(n int) {
	var f0 int // bắt đầu bằng 0
	var f1 = 1 // và bắt đầu bằng 1
	var fn = 1 //
	for i := 0; i < n; i++ {
		if n < 0 { // n luôn dương
			n = -1
		} else if n == 0 || n == 1 { // n không bằng 0 hoặc 1
			n = n
		} else {
			fmt.Print(f0, f1, fn, " ")
			for i = 2; i < n; i++ {

				f0 = f1
				f1 = fn
				fn = f0 + f1
				if fn < n {
					fmt.Print(fn, " ")
				}
			}
		}
	}
}
