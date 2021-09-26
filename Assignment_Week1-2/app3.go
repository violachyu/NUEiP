package main

import (
	"fmt"
)

func main() {
	n := 3
	add(&n, 2)
	fmt.Println(n) // 印出 5
	add(&n, 5)
	fmt.Println(n) // 印出 10
}

// 完成一個 add 函式，讓主程式能印出預期的結果
func add(num1 *int, num2 int) {
	// var temp int
	// temp = *num1 // save num1 data value to temp
	// num1 = temp + num2
	*num1 += num2
}

/* Reference */
// https://www.tutorialspoint.com/go/go_passing_pointers_to_functions.htm
