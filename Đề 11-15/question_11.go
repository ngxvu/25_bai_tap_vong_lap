// Kiểm tra n có phải là số nguyên tố hay ko
// kiem tra == 0 ( so sánh bằng ), ví dụ 5==5

// n=4 &&  i=2 3

package main

import "fmt"

func main() {
	var n int = 4
	kiemtra := 1
	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			kiemtra = 0
			break
		}
	}
	if kiemtra == 1 {
		fmt.Println("n là số nguyên tố")
	} else if kiemtra == 0 {
		fmt.Println("n ko phải là số nguyên tố")
	}
}
