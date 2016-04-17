package main

import "fmt"

type Task struct {
	ID     int
	Detail string
	done   bool
}

// 型内にString()というメソッドを定義
func (task Task) String() string {
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

func (task *Task) Finish() {
	task.done = true
}

func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
	return task
}

func main() {
	task := NewTask(1, "buy the milk")
	task.Finish()
	fmt.Printf("%s", task)
}
