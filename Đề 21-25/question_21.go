// Nhập một số nguyên dương n. Viết chương trình kiểm tra số n có phải là số hoàn hảo hay không?

package main

import "fmt"

func main() {
	var n int = 33
	var i, s int
	if n <= 0 {
		fmt.Println("N phai lon hon 0")
	} else {
		for i = 1; i <= n-1; i++ {
			if n%i == 0 {
				s += i
			}
		}
		if s == n {
			fmt.Print(n, " la so hoan hao")
		} else {
			fmt.Print(n, " khong phai la so hoan hao")
		}
	}
}
