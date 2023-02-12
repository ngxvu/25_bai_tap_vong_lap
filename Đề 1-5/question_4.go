// Viết chương trình nhập một số nguyên "x", tìm bội số của số đó với các số từ 1 đến 20, sau đó in kết quả ra màn hình.

package main

import (
	"fmt"
)

func main() {
	var x int = 5
	fmt.Println("Nhập một số nguyên, tìm bội số của số đó với các số từ 1 đến 20.")
	for i := 1; i <= 20; i++ {
		fmt.Println(i, "*", x, "=", i*x)
	}
}
