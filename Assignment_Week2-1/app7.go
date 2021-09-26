package main

import (
	"fmt"
)

// 請設計 Man 結構，定義姓名和年齡，並滿足主程式的需求
func main() {
	m := Man{"John", 30}
	print(m) // 印出：I’m John, 30 years old.
	s := "Hello"
	print(s) // 印出：String is Hello
	f := 3.5
	print(f) // 印出：Float is 3.5
}

// 請運用空白介面當參數，建立 print 方法，滿足主程式需求
func print(val interface{}) {
	switch v := val.(type) {
	case string:
		fmt.Printf("String is %v", v)
	case float32:
		fmt.Printf("Float is %v", v)
	case Man:
		fmt.Printf("I'm %s, %d years old.", v.Name, v.Age)
	}
}

type Man struct {
	Name string
	Age  int
}

//Ans. empty interface

/*
reference: https://tour.golang.org/methods/14

*/
