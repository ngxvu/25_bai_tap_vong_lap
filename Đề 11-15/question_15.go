// Nhập số nguyên dương n bất kỳ. Viết chương trình vẽ tam giác "*" có chiều cao là n hàng.

/*	input 4 >> output
****
***
**
*
 */

package main

import "fmt"

func main() {

	var n int = 4
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
