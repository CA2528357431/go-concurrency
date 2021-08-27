// 协程保护


package main

import (
	"fmt"
	"sync"
)

var (
	res = make(map[int]int)
	lock sync.Mutex
	check = make(chan int,10)


)

func main() {


	for i:=0;i<10;i++{
		go do(i)
	}

	// 每个协程结束都会向check里插入一个数字
	// 在此我们要求从check弹出10个数据，否则阻塞
	for i:=0;i<10;i++{
		<-check
	}

	for key,value:=range res{
		fmt.Println(key,value)
	}
	fmt.Println(len(res))



}
func do(x int)  {
	for ;x<100;x+=10{
		lock.Lock()
		res[x] = x*x*x
		lock.Unlock()
	}
	check<-1

}
