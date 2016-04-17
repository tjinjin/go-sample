package main

import "fmt"

type Task struct {
	ID     int
	Detail string
	done   bool
}

func main() {
	var task Task = Task{
		ID:     1,
		Detail: "buy the milk",
		done:   true,
	}
	fmt.Println(task.ID)
	fmt.Println(task.Detail)
	fmt.Println(task.done)

	// 順番に値を格納する

	var test Task = Task{
		2, "buy the apple", true,
	}
	fmt.Println(test.ID)
	fmt.Println(test.Detail)
	fmt.Println(test.done)

	// 特に値を指定しない場合はゼロ値で初期化
	aaa := Task{}
	fmt.Println(aaa.ID)     // 0
	fmt.Println(aaa.Detail) // ""
	fmt.Println(aaa.done)   // false
}
