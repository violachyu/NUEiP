package main

import "fmt"

func main() {
	// 定義變數 ch，建立並存放適當的 channel，讓以下程式能印出 true 和 false
	ch := make(chan bool)
	ch <- true
	ch <- false
	ch <- true
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Doing
