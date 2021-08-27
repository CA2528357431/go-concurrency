// 协程错误

// 一个协程的错误就能使整个程序停转

package main

import (
	"errors"
	"fmt"
	"time"
)

var (
	exit = make(chan bool,1)
)

func main() {

	go alpha()
	go beta()

	time.Sleep(time.Second*3)

	fmt.Println("finish")

}


func beta()  {
	for i:=0;i<5;i++{
		fmt.Println(i,"b")
	}
	exit<-true
}


func alpha()  {

	defer check()

	for i:=0;i<5;i++{
		if i==3{
			err := errors.New("three")
			panic(err)
		}
		fmt.Println(i,"a")

	}

}

func check()  {
	err := recover()
	if err!=nil{
		fmt.Println("get error")
	}
}