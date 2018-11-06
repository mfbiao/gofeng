package gofeng

import (
	"fmt"
	"gofeng/pipeline"
	"os"
	"bufio"
)

func main(){
	const filename = "./small.in"
	const count = 10000000
	file,err := os.Create(filename)
	defer file.Close()
	if err != nil{
		panic(err)
	}
	p := pipeline.RandomSource(count)
	//pipeline.WriteSink(file,p)
    //用bufio
	write := bufio.NewWriter(file)
	pipeline.WriteSink(write,p)
	write.Flush()

	file,err = os.Open(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	//p = pipeline.ReaderSource(file)
	p = pipeline.ReaderSource(bufio.NewReader(file),-1)
	i := 0
	for v := range p {
		i++
		if i <100 {
			fmt.Println(v)
		}
	}

}

func mergeDemo(){

	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3,5,23,4,8)),
		pipeline.InMemSort(pipeline.ArraySource(10,24,5,8,24)))
	for v := range p{
		fmt.Println(v)

	}

	//p := pipeline.ArraySource(3,6,4,333,543,3443,322)

	//for {
	//	//if num,ok := <-p;ok{
    	//	//fmt.Println(num)
	//	//}else{
	//	//	break
	//	//}
	//
	//
	//	for v := range p{
	//		fmt.Println(v)
	//	}
	//}
}

