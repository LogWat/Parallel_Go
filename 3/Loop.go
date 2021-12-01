package main

import (
	"fmt"
	"sync"
)

// ゴルーチンの意外性
// 以下の場合，ゴルーチンが開始する前にループが終了する場合がほとんど
// ゴルーチンはスコープ外からsalutationを参照する このとき，salutationは参照が
// 残っているため，ヒープに移っている
func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() { // 目的通りの動作をさせるならばここに引数を与える
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}
