package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func (u *User) FullName() string {
	fullname := fmt.Sprintf("%s %s",
		u.FirstName, u.LastName)
	return fullname
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}

type Task struct {
	ID     int
	Detail string
	done   bool
	*User  // Userを埋め込む
}

func NewTask(id int, detail, firstName, lastName string) *Task {

	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
		User:   NewUser(firstName, lastName),
	}
	return task
}

func main() {
	task := NewTask(1, "buy the milk", "Jxck", "Daniel")
	fmt.Println(task.FirstName)
	fmt.Println(task.LastName)
	fmt.Println(task.FullName())
	fmt.Println(task.User)
}
