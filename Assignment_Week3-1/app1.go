package main

import (
	"fmt"
	"time"
)

func main() {
	go print("1") // 修改這行程式，啟動 goroutine，使得程式有機會先印出 2，然後才印出 1
	print("2")
	time.Sleep(100 * time.Millisecond)
}
func print(msg string) {
	fmt.Println(msg)
}

// Done.
