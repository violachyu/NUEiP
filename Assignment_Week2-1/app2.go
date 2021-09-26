package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	print(s[1:4]) // 請使用切片調整傳入的參數，讓函式能印出：len=3 cap=5 [3 5 7]
	print(s[5:6]) // 請使用切片調整傳入的參數，讓函式能印出：len=1 cap=0 [13]
}
func print(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Ans Done, 是否太過大幅更動？
