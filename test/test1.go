package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}
	wg.Wait()
	//for i:=0;i<10;i++{
	//	<-c
	//}
	//	time.Sleep(0.1*time.Second);
	//	<-c
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	time.Sleep(1 * time.Second)
	fmt.Println(index, a)
	//if index ==9 {
	//c  <-true
	wg.Done()
	//}
}
