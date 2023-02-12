package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Nhập số n: ")
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())

	for i := 1; i <= 20; i++ {

		fmt.Print("Bội số của n là: ", input*i)

		fmt.Println()
	}
	fmt.Println()

}
