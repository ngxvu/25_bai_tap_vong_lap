// Nhập số nguyên dương n bất kỳ, viết chương trình in n phần tử đầu tiên của dãy Fibonacci.

/* input: 9 >> output: 0 1 1 2 3 5 8 13 21 ( 9 số )

Nhập giá trị cho biến n
Kiểm tra n:
   Nếu n <=2
      Thông báo nhập lại giá trị lớn hơn 2
   Nếu n >2
      In f1 và f2
      Cho biến i chạy từ 3 đến n
         Tính f=f1+f2
         In f ra
         Gán f1=f2
         Gán f2=f

*/

/* package main

import "fmt"

func main() {
	var n int = 9
	f1 := 1
	f2 := 1

	if n <= 2 {
		fmt.Print("n khong hop le")
	}
	if n > 2 {
		fmt.Print(f1, f2)
	}
}
*/

package main

import "fmt"

func main() {
	Solution(20)
}

func Solution(n int) {
	var f0 int // bắt đầu bằng 0
	var f1 = 1 // và bắt đầu bằng 1
	var fn = 1 //
	for i := 0; i < n; i++ {
		if n < 0 { // n luôn dương
			n = -1
		} else if n == 0 || n == 1 { // n không bằng 0 hoặc 1
			n = n
		} else {
			fmt.Print(f0, f1, fn)
			for i = 2; i < n; i++ { // tính số fibo
				f0 = f1
				f1 = fn
				fn = f0 + f1
				fmt.Print(" ", fn)
			}
		}
	}
}
