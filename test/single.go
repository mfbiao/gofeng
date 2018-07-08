package main

import (
	"fmt"
	"sync"
)

type single struct {
	Name string
}

var m *single
var lock sync.Mutex
var once sync.Once

// 最简单的一种，懒汉模式
//func GetInstance() *single {
//	if m == nil{
//		m = &single{}
//	}
//	return m
//}
//双重锁机制：这次我们用了两个判断，而且我们将同步锁放在了条件判断之后，这样做就避免了每次调用都加锁，
// 提高了代码的执行效率。理论上写到这里已经是很完美的单例模式了
//func GetInstance()*single{
//	if m == nil{
//		lock.Lock()
//		defer lock.Unlock()
//		if m == nil{
//			m = &single{}
//		}
//	}
//	return m
//}
//Once.Do方法的参数是一个函数，这里我们给的是一个匿名函数，
// 在这个函数中我们做的工作很简单，就是去赋值m变量，而且go能保证这个函数中的代码仅仅执行一次！
func GetInstance() *single {
	once.Do(func() {
		m = &single{}
	})
	return m
}

func main() {
	a := GetInstance()
	a.Name = "a"
	b := GetInstance()
	b.Name = "b"
	fmt.Println(&a.Name, a)
	fmt.Println(&b.Name, b)
	fmt.Printf("%p %T\n", a, a)
	fmt.Printf("%p %T\n", b, b)
}
