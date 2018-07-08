package main

import (
	"fmt"
	"runtime"
	//	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
	//	time.Sleep(0.1*time.Second);
	//	<-c
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}

	fmt.Println(index, a)
	//if index ==9 {
	c <- true
	//}
}
