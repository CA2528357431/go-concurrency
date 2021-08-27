// 只读只写

package main

import "fmt"

func write(channel chan<- int)  {
	channel<-1
	channel<-2
	channel<-3
	channel<-4
	channel<-5
	// x := <-channel报错
	// 本函数中充当只写
}



func read(channel <-chan int)  {
	x := <-channel
	fmt.Println(x)
}


func main() {

	channel := make(chan int,5)

	write(channel)
	read(channel)






}
