// channel

// 本身线程安全
// FIFO

package main

import (
	"fmt"
)


func main() {

	var mychannel chan int
	mychannel = make(chan int,5)
	// 设置channel容量
	// 等效于 mychannel := make(chan map[int]int,3)

	mychannel<-1
	x := 9
	mychannel<-x
	// 放入量

	mychannel<-3
	mychannel<-8


	fmt.Println(len(mychannel),cap(mychannel))
	// 长度和容量


	close(mychannel)
	// 从此不能在放入，但可以取出

	x,ok:=<-mychannel
	if !ok{
		fmt.Println(ok)
	}
	fmt.Println(x)
	// 取值
	// 若未关闭，在已经空的chan里取值，会阻塞直到加入数据
	// 同理如果已经满了还在插入，则阻塞到取出数据

	for value := range mychannel{
		fmt.Println(value,len(mychannel))
	}
	// 关闭后才能遍历
	// 遍历获值的过程其实是chan弹出数值的过程






}
