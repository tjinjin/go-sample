package main

import "fmt"

type Task struct {
	ID     int
	Detail string
	done   bool
}

// 構造体への変更が反映されない
func Finish(task Task) {
	task.done = true
}

func Finish_pointer(task *Task) {
	task.done = true
}

func main() {
	task := Task{done: false}
	Finish(task)
	fmt.Println(task.done) // falseのまま

	test := &Task{done: false}
	Finish_pointer(test)
	fmt.Println(test.done) // true
}
