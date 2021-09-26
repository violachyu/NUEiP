package main

import "fmt"

func main() {
	// 定義變數 ch，建立並存放適當的 channel，滿足以下程式要求
	ch := make(chan int)
	max := 10
	go calculate(max, ch)
	fmt.Println(<-ch) // 從 ch 接收資料，並印出來
}

// 建立 calculate 函式，使用 goroutine 執行，計算 1+2+...+max 的總和
// 將計算結果透過 channel 傳遞到主程式中

func calculate(max int, ch chan int) {
	result := 0
	for i := 0; i < max; i++ {
		result += i
	}
	ch <- result
}

// Done.
