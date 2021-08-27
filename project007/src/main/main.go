// select

package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int,5)
	c2 := make(chan string,10)
	c1<-1
	c1<-8
	c1<-3
	c1<-5
	c1<-4
	c2<-"hello,go"
	c2<-"hello,go"
	c2<-"hello,go"
	c2<-"hello,go"
	c2<-"hello,go"
	c2<-"hello,go"
	c2<-"hello,go"
	c2<-"hello,go"



	// go add(c2)
	// 所谓有无数值只是指那一瞬间有无数据，可能还有数据在添加

	for {
		time.Sleep(time.Millisecond*100)
		select {
		case i:=<-c1:
			fmt.Println(i,"int")
		case s:=<-c2:
			fmt.Println(s,"string")
		default:
			fmt.Println("no")
			break
		}

	}
	// 只会从可选的channel中得到数据，
	// 否则：若有default则执行；若无则堵塞
}

func add(x chan string)  {
	for i:=0;i<10;i++ {
		time.Sleep(time.Millisecond*500)
		x<-"hello,world"
	}
}