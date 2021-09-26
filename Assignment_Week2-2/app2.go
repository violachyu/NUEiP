package main

import "fmt"

func main() {
	// 在此，請適當的定義 data 變數，存放地圖資料，讓以下程式能正常運作
	data := make(map[string]string)
	data["apple"] = "蘋果"
	data["dog"] = "狗狗"
	data["cat"] = "貓貓"
	// 在此，請要求使用者透過終端機輸入英文字，利用 data 進行對應中文字的查詢
	// 若輸入的英文字包含在 data 的 key 中，則印出對應的中文；若無，印出沒有查到
	fmt.Println("Enter Random Word: ")
	var word string

	fmt.Scanln(&word)
	// i := 0; i < len(data); i++
	for key, _ := range data {
		value, p := data[word]
		if p == true && key == word {
			fmt.Println(value)
			return
		} else {
			fmt.Println("Word not found")
			return
		}
	}
}

// Done.
/*
-- initialization of Map
-- Error: panic: assignment to entry in nil map
	https://github.com/kevinyan815/gocookbook/issues/7
-- take input from user: https://www.geeksforgeeks.org/how-to-take-input-from-the-user-in-golang/
*/
