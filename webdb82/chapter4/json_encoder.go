// []byte経由せずに直接jsonとして書き込む
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"-"`
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"`
	memo    string
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

	// ファイルを開く
	file, err := os.Create("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// エンコーダの取得
	encoder := json.NewEncoder(file)
	fmt.Println(file)
	fmt.Println(encoder)

	// JSONエンコードしたデータの書き込み
	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person)
}
