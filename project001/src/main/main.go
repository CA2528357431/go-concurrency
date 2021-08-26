// 协程入门

// 当主线程结束，其他goroutine无论是否结束都直接结束

package main

import (
	"fmt"
	"time"
)

func main() {
	// 另开协程，从此分离于主线程
	go appenddo()

	maindo()

}

func maindo()  {
	for i:=0;i<5;i++{
		time.Sleep(500*time.Millisecond)
		fmt.Println("main",i)
	}
}

func appenddo()  {
	for i:=0;i<10;i++{
		time.Sleep(500*time.Millisecond)
		fmt.Println("append",i)
	}
}