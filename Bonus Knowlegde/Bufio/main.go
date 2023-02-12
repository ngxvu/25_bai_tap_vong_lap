package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solution(songuyen int) int {
	scanner := bufio.Scanner(os.Stdin)
	fmt.Printf("Nhập số nguyên: ")
	scanner.Scan()
	input := strconv.ParseInt(scanner.Text())
	fmt.Printf("Dạng số đảo ngược là: %d", )