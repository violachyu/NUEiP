package main

func main() {
	// 定義變數 ch1 和 ch2，建立並存放適當的 channel，滿足以下程式要求
	ch1 := make(chan int)
	ch2 := make(chan int)
	max := 10
	go calculate(max, ch1)
	max = 11
	go calculate(max, ch2)
	// 以上有兩個 goroutines 同時執行
	// 在此使用 select 語句，判斷哪個 channel 最先取得資料，印出最先接收到的總和
	// 程式每次執行時，結果可能是 55 或 66
}

// 建立 calculate 函式，使用 goroutine 執行，計算 1+2+...+max 的總和
// 將計算結果透過 channel 傳遞到主程式中

// Doing
