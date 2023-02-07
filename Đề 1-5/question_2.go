// Viết chương trình in các số từ 1 đến 1000 theo thứ tự giảm dần.

package main

import "fmt"

func main() {
	fmt.Println("In các số từ 1 đến 1000 theo thứ tự giảm dần.")
	for x := 1000; x >= 1; x-- {
		fmt.Println(x)
	}
}
