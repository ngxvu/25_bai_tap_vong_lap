// Viết chương trình nhập một câu bất kỳ, đếm số từ và ký tự trong câu đó, và in kết quả ra màn hình.

/*

--- input 					 output ---
-the gioi di dong ---------- So tu: 4
							 So ky tu: 3

Giai phap:

	Nhap:"The gioi di dong" >> Loop qua tung chu

*/

package main

import "fmt"

func main() {
	Output, Output2, Output3 := Giaiphap("the gioi di dong")
	fmt.Println(Output)
}

func Giaiphap(S string) (int, int, string) {
	demtu := 0
	demkytu := 0
	for _, v := range S {
		if v == 32 {
			demkytu++
		}
		if v != 32 {
			demtu++
		}
	}
	return demtu, demkytu, "hello"
}
