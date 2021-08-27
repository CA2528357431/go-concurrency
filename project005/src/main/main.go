// 管道通信

package main

import (
	"fmt"
	"math"
)

var (
	dock = make(chan int,50)
	res = make(chan int,200)
	fin = make(chan bool,10)
)

func check()  {
	for {
		x, ok := <-dock
		if !ok{
			break
		}
		limit := int(math.Pow(float64(x), 0.5))
		judge := true
		for r := 2; r <= limit; r++ {
			if x%r == 0 {
				judge = false
				break
			}
		}
		if judge {
			res <- x
		}

	}
	fin <- true
}



func add()  {
	for i := 2; i < 1000; i++ {
		dock<-i
	}
	close(dock)
}


func main() {

	go add()

	for i:=0;i<10;i++{
		go check()
	}


	for i:=0;i<10;i++{
		<-fin
	}
	close(res)


	resli := make([]int,0)
	for  {
		x, ok := <-res
		if !ok{
			break
		}
		resli = append(resli,x)
	}

	for _,v := range resli{
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println(len(resli))

}