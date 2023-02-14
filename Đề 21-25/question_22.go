// Nhập x là số thực, y là số tự nhiên. Hãy tính xy.

package main

import "fmt"

func main() {
	var x int = 5
	var y int = 4
	var luy_thua = 1
	for i := 1; i <= y; i++ {
		luy_thua *= x
	}
	fmt.Print(luy_thua)
}
