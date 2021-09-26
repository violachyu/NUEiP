package main

import "fmt"

func main() {
	validate(50)     // Too Few
	validate(200)    // OK
	validate(20000)  // OK
	validate(200000) // Too Much
}
func validate(money int) {
	// 請運用判斷式，完成以下判斷
	// 如果輸入的 money 不到 100，印出：Too Few
	// 如果輸入的 money 在 100 ~ 100000 之間，印出：OK
	// 如果輸入的 money 超過 100000，印出：Too Much
	switch {
	case money < 100:
		fmt.Println("Too Few")
	case money > 100 && money < 100000:
		fmt.Println("OK")
	default:
		fmt.Println("Too Much")
	}
}
