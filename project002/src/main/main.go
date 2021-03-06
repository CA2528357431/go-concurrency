// 资源保护


package main

import (
	"fmt"
	"sync"
)

var (
	res = make(map[int]int)
	// 互斥锁
	lock sync.Mutex
)

func main() {


	for i:=0;i<10;i++{
		go do(i)
	}


	// 防止读取的时候被改写
	lock.Lock()
	for key,value:=range res{
		fmt.Println(key,value)
	}
	fmt.Println(len(res))
	lock.Unlock()


}
func do(x int)  {
	for ;x<100;x+=10{
		lock.Lock()
		res[x] = x*x*x
		lock.Unlock()
		// 对map进行保护
	}
}


// 我们会发现大量数据尚未被写入
// 因为我们尚未进行协程保护