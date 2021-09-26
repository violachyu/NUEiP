package main

import "fmt"

// 請設計 Man 和 Bird 結構以及對應的 Talk 方法，另外設計 Talkable 介面，滿足主程式需求
// Bird 結構包含姓名、年齡和飛行距離的資訊
func main() {
	var t Talkable = Man{"John", 30}
	t.Talk() // 印出：I’m John, 30 years old.
	t = Bird{"KiKi", 2, 10}
	t.Talk() // 印出：I’m bird KiKi, 2 years old. Fly 10 miles a day.
}

type Talkable interface {
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

func (m Man) Talk() {
	fmt.Printf("I’m %s, %d years old.", m.Name, m.Age)
}

func (b Bird) Talk() {
	fmt.Printf("I’m bird %s, %d years old. Fly %d miles a day.", b.Name, b.Age, b.Distance)
}

// Ans. 學會使用interface
