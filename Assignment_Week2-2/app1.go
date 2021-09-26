package main

func main() {
	// 在此，請適當的定義 data 變數，存放雙層切片的資料，讓以下程式能正常運作
	var data [2][3]int
	data[0][0] = 10
	data[0][1] = 5
	data[0][2] = 1
	data[1][0] = 3
	data[1][1] = 8
	data[1][2] = 4
	// 在此，請運用雙層 for 迴圈，計算以上六個數字的總和，並印出結果
	result := 0
	// fmt.Println(len(data))
	for i := 0; i < len(data); i++ {
		for _, columnValue := range data[i] {
			result += columnValue
		}
	}
	println(result)
}

// Done.
// extQ. 一定要指定長度？
