package concurrency

import (
	"fmt"
	"time"
	"runtime"
)

//M-P-G

func sayShort(s string) {
	for i := 0; i < 5; i++ {
		//runtime.Gosched()
		fmt.Println(s)
	}
}

func DemoFirstGoRtn() {
	runtime.GOMAXPROCS(4)

	go sayShort("world")
	sayShort("hello")

	time.Sleep(1)
}
