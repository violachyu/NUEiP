package main

import "fmt"

func main() {
	n := 3
	n = add(n, 2)
	fmt.Println(n) // 印出 5
	n = add(n, 5)
	fmt.Println(n) // 印出 10
}

// 完成一個 add 函式，讓主程式能印出預期的結果
func add(num1 int, num2 int) int {
	result := num1 + num2
	return result
}
