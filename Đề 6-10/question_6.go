// Nhập số nguyên n bất kỳ. Viết chương trình in các số chẵn từ 1 đến n.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.Reader{os.Stdin}
	fmt.Println("nhap so n: ")
	scanner.Scan()
	for i := 1; i <= x; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}
