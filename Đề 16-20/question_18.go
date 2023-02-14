// Viết chương trình tính kết quả biểu thức S= 1+1/2^3+1/3^3+...+1/n^3 (làm tròn 5 chữ số thập phân).

package main

import "fmt"

func main() {
	n := 6
	var S float64 = 0
	for i := 1; i <= n; i++ {
		S += 1 / float64(i*i*i)
	}
	fmt.Print(S)

}
