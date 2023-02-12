// Viết chương trình nhập một số nguyên n và in kết quả ra màn hình dưới dạng số đảo ngược (về thứ tự) của số nguyên vừa nhập đó.

/*
123456 > 0
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solution(songuyen int) int {
	sodu := 0
	for songuyen > 0 {
		chialaydu := songuyen % 10
		sodu = (sodu * 10) + chialaydu
		songuyen /= 10
	}
	return sodu
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Nhập số nguyên của bạn: ")
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())
	a := Solution(input)
	fmt.Println("Số đảo ngược là: ", a)

}
