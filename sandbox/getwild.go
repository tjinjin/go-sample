//see more: http://qiita.com/takochuu/items/20f639194ac885fb3e1e
package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"sync"
	"time"
)

var (
	getWildWords = []string{"G", "e", "t", "W", "i", "l", "d"}
	g            string
)

const (
	textGetWild = `
  ____      _    __        ___ _     _
 / ___| ___| |_  \ \      / (_) | __| |
| |  _ / _ \ __|  \ \ /\ / /| | |/ _  |
| |_| |  __/ |_    \ V  V / | | | (_| |
 \____|\___|\__|    \_/\_/  |_|_|\__,_|
`
)

// 比較用に答えを作っておく
func init() {
	for _, v := range getWildWords {
		g += v
	}
}

func main() {
	// goroutinの終了タイミングをあわせる
	wg := &sync.WaitGroup{}
	// string型のチャネル
	c := make(chan string)

	for {
		w := ""
		for i := 0; i < 7; i++ {
			// wait groupを追加
			wg.Add(1)
			go getwildAndTough(wg, c)
			w += <-c
		}
		// WaitGroupの完了待ちここではforが7回回るので7つ待つ
		wg.Wait()
		if isGetWild(w) {
			fmt.Println(textGetWild)
			exec.Command("say", "-v", "Ava", w).Run()
			break
		}
		w = ""
	}
	close(c)
}

func getwildAndTough(wg *sync.WaitGroup, c chan string) {
	// WaitGroupが完了したことを通知
	// deferは関数終了時に実行したい書くときに使う
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	// Intn(n) で0からnまでの数字からランダムで数字を選ぶ
	c <- getWildWords[rand.Intn(len(getWildWords))]
}

func isGetWild(w string) bool {
	if g == w {
		return true
	}
	return false
}
