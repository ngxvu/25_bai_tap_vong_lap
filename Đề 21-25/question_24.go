// Nhập vào hai số nguyên dương a và b. Tìm tổng của tất cả các số lẻ và chẵn nằm giữa hai số đó.

package main

import "fmt"

func main() {
	var a int = 5
	var b int = 9

	if a == b {
		fmt.Println("a khong duoc bang b")
	}
	tong_so_chan := 0
	tong_so_le := 0
	if a > b {
		for i := a; i >= b; i-- {
			if i%2 == 0 {
				tong_so_chan += i
			} else {
				tong_so_le += i
			}
		}
		fmt.Println("Tổng số lẻ", tong_so_le)
		fmt.Println("Tổng số chẵn", tong_so_chan)
	} else {
		for i := b; i >= a; i-- {
			if i%2 == 0 {
				tong_so_chan += i
			} else {
				tong_so_le += i
			}
		}
		fmt.Println("Tổng số lẻ", tong_so_le)
		fmt.Println("Tổng số chẵn", tong_so_chan)
	}

}
