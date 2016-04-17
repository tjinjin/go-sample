// データとして出力させたいものを調整する

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	ID      int    `json:"id"`                // idという名前で格納
	Name    string `json:"name"`              // nameという名前で格納
	Email   string `json:"-"`                 // JSONに格納しない
	Age     int    `json:"age"`               // ageという名前で格納
	Address string `json:"address,omitempty"` // 値が空なら無視
	memo    string
	// Address string `json:",string"` 値をJsonとして格納
}

func main() {
	person := &Person{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "golang lover",
	}
	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b)) // 文字列に変換
	// {"id":1,"name":"Gopher","age":5} 出力結果
	// {"ID":1,"Name":"Gopher","Email":"gopher@example.org","Age":5,"Address":""}

}
