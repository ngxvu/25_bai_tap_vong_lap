// Nhập một số n nguyên dương bất kỳ, viết chương trình in các số nguyên tố từ 0 đến n bằng vòng lặp FOR.

// *1 khong phai la so nguyen to
// * so nguyen to la so chia het cho 1 va chinh no
// * neu so chia het cho 1 so thu 3 tro di do ko phai la so nguyen to

package main

import "fmt"

func main() {
	// khai bao bien
	n := 10
	fmt.Println("Cac so nguyen to can tim la: ")
	var kiemtra int
	for i := 1; i <= n; i++ {
		for j := 2; j <= i-1; j++ {
			if i%j == 0 {
				kiemtra = 0
			}
		}
		if kiemtra == 1 {
			fmt.Printf("%d ", i)
		} else {
			kiemtra = 1
		}
	}
}
