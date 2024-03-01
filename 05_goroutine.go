package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world") //開一個新的 Goroutines 執行
	say("hello")    //當前 Goroutines 執行
}

//goroutine
//就是協程 (Coroutine)，但是它比執行緒更小
//可同時執行成千上萬個併發任務。goroutine 比 thread 更易用、更高效、更輕便。

//透過 Go 的 runtime 管理的一個執行緒管理器。
//goroutine 透過 go 關鍵字實現了，其實就是一個普通的函式。


//go 關鍵字很方便的就實現了併發程式設計。 
//上面的多個 goroutine 執行在同一個程序裡面，共享記憶體資料，
//不過設計上我們要遵循：不要透過共享來通訊，而要透過通訊來共享。
//runtime.Gosched()表示讓 CPU 把時間片讓給別人，下次某個時候繼續恢復執行該 goroutine。