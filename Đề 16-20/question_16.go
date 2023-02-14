// Viết chương trình nhập n số dương. Chương trình sẽ kết thúc nếu một trong các số đó là số âm.

package main

import (
	"fmt"
)

func main() {
	var n int = 0
	for i := 0; i <= n; i++ {
		if n <= -15 {
			fmt.Print("n phải là số dương ")
		} else {
			fmt.Print(i, " ")
		}
	}
}
