package main

import (
	"fmt"
	"time"
)

// 請在此定義變數 mu，儲存 sync.Mutex 物件，並在下方程式中運用以解決資料混亂問題
var count int = 0

func main() {
	go add(10000)
	go add(20000)
	time.Sleep(time.Second)
	fmt.Println(count) // 這行應該要穩定的印出 30000
}
func add(amount int) {
	for n := 0; n < amount; n++ {
		count++
	}
}
