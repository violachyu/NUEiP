package main

import (
	"fmt"

	"github.com/drgrib/iter"
)

func main() {
	// 定義變數 ch，建立並存放適當的 channel，滿足以下程式要求
	ch := make(chan int)
	max := 10
	go iterate(max, ch)
	// 運用 for 迴圈搭配 range 關鍵字，持續從 ch 接收資料，並印 1, 2, 3, …, max
}

// 建立 iterate 函式，使用 goroutine 執行
// 運用迴圈將 1, 2, 3, …, max 的資料持續傳遞回主程式，最後運用 close 方法關閉 channel
func iterate(max int, ch chan int) {
	for i := range iter.N(max) {
		ch <- i
		fmt.Println(ch)
	}
	close(ch)
}

// ??
/*
reference:
-- [range int] https://stackoverflow.com/questions/21950244/is-there-a-way-to-iterate-over-a-range-of-integers
-- [closing channels] https://gobyexample.com/closing-channels

*/
