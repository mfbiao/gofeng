package main

import "fmt"

type stack struct { // ← 栈不应该被导出
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
}
func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

//像下面这样使用
func main() {
	//var s stack
	s := new(stack)
	s.push(25)
	s.push(14)
	fmt.Printf("stack %v\n", s)
}
