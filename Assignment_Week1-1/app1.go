package main

import (
	"fmt"
)

func main() {
	n1 := 2
	n2 := 3
	newN1 := float32(n1)
	newN2 := float32(n2)
	result := newN1 / newN2
	fmt.Println(result)
	// 請輸出 n1 除以 n2 的結果，結果必須是小數 ( 浮點數 ) 的資料型態
}
