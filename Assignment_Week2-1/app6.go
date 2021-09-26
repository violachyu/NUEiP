package main

import "fmt"

// 承上，設計 Animal 介面，滿足主程式需求
func main() {
	m := Man{"John", 30}
	// 修正以下參數傳遞的方式，需印出：I’m John, 31 years old.
	birthdayTalk(&m)
	b := Bird{"KiKi", 2, 5}
	// 修正以下參數傳遞的方式，需印出：I’m bird KiKi, 3 years old. Fly 5 miles a day.
	birthdayTalk(&b)
}
func birthdayTalk(a Animal) {
	a.Grow(1)
	a.Talk()
}

type Animal interface {
	Grow(addAge int)
	Talk()
}

type Man struct {
	Name string
	Age  int
}

type Bird struct {
	Name     string
	Age      int
	Distance int
}

func (m *Man) Grow(addAge int) {
	m.Age += addAge
}

func (b *Bird) Grow(addAge int) {
	b.Age += addAge
}

func (m Man) Talk() {
	fmt.Printf("I’m %s, %d years old.", m.Name, m.Age)
}

func (b Bird) Talk() {
	fmt.Printf("I’m bird %s, %d years old. Fly %d miles a day.", b.Name, b.Age, b.Distance)
}

//Done.
/*
-- [instantiate an interface]: using struct (type Man, type Bird, func Grow, func Talk)
-- [taking interface as a param]: birthdayTalk()
-- [call by reference(pointer)] to modify the original value
*/
