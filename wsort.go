package gofeng

import (
	"os"
	"gofeng/pipeline"
	"bufio"
	"fmt"
)

func main(){
   p:= createPipeline("./small.in",80000000,4)
   writeToFile(p,"./small.out")
   printFile("./small.out")

}

func printFile(filename string){
	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	p := pipeline.ReaderSource(file,-1)
	i := 0
	for v := range p{
		i++
		if i <= 10 {
			fmt.Println(v)
		}
	}
}

func writeToFile(p <- chan int,filename string){
   file,err	:= os.Create(filename)
   if err != nil {
    	panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriteSink(writer,p)

}

func createPipeline(filename string,fileSize,chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	sortResults := []<-chan int{}
	pipeline.Init()
	for i:=0;i<chunkCount;i++{
		file,err := os.Open(filename)
		if err != nil{
			panic(err)
		}
        file.Seek(int64(i*chunkSize),0)
        source := pipeline.ReaderSource(bufio.NewReader(file),chunkSize)
		sortResults = append(sortResults,pipeline.InMemSort(source))
    }
    return  pipeline.MergeN(sortResults...)
}