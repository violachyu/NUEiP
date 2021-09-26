package main

func main() {
	var n1 int = 2
	var n2 int = 3
	divide(n1, n2)
}

// 請建立 divide 函式，輸出 n1 除以 n2 的結果，結果必須是小數 ( 浮點數 ) 的資料型態

func divide(n1 int, n2 int) float32 {
	newN1 := float32(n1)
	newN2 := float32(n2)
	result := newN1 / newN2
	return result
}
