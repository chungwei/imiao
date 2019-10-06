package main

import (
	"fmt"
	"math"
	"sync"
)

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}
func main() {
	//wg := new(sync.WaitGroup)
	//wg.Add(2)
	//for i := 0; i < 2; i++ {
	//	go func(id int) {
	//		defer wg.Done()
	//		sum(id)
	//	}(i)
	//}
	//wg.Wait()

	var wg sync.WaitGroup

	wg.Add(2) // 因为有两个动作，所以增加2个计数
	go func() {
		fmt.Println("Goroutine 1")
		wg.Done() // 操作完成，减少一个计数
	}()

	go func() {
		fmt.Println("Goroutine 2")
		wg.Done() // 操作完成，减少一个计数
	}()

	wg.Wait() // 等待，直到计数为0

	s := 1<<20 - 1
	fmt.Println(`----`, s)

	c := make(chan int, s)
	for i := 0; i < s; i++ {

		go func(i int) {
			//fmt.Println(i)
			c <- i
		}(i)
	}

	for i := 0; i < s; i++ {
		<-c
		//fmt.Println(t)
	}
}
