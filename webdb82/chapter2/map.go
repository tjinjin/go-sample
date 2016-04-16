package main

import "fmt"

func main() {
	//var month map[int]string = map[int]string{}
	//
	//month[1] = "January"
	//month[2] = "February"
	//fmt.Println(month)

	month := map[int]string{
		1: "January",
		2: "February",
	}

	fmt.Println(month)

	// mapの操作
	jan := month[1]
	fmt.Println(jan)

	//戻り値だけ受け取る
	_, ok := month[1]
	if ok {
		fmt.Println("aaa")
	}

	//データを削除する
	delete(month, 1)
	fmt.Println(month) // map[1:January]

	// key-valueを受け取りながら処理をすすめる。マップの順番は保証されない
	for key, value := range month {
		fmt.Printf("%d %s\n", key, value)
	}

}
