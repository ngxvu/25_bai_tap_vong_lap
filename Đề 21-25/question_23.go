// Nhập một số nguyên dương n bất kỳ. Viết chương trình kiểm tra số n có phải là số nguyên tố không?

package main

import "fmt"

func main() {
	var n int = 49
	kiemtra := 1
	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			kiemtra = 0
			break
		}
	}
	if kiemtra == 1 {
		fmt.Println(n, "là số nguyên tố")
	} else if kiemtra == 0 {
		fmt.Println(n, "không phải là số nguyên tố")
	}
}
