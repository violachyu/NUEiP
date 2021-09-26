package main

import (
	"fmt"
	"time"
)

var count int = 0

func main() {
	go add(10000)
	go add(20000)
	time.Sleep(time.Second)
	fmt.Println(count) // 這行應該要印出多少，試著多執行幾次，請解釋此現象
}
func add(amount int) {
	for n := 0; n < amount; n++ {
		count++
	}
}
