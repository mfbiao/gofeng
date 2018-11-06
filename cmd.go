package main

import (
	"runtime"
	"os/exec"
	"fmt"
)

func main(){
	cos := runtime.GOOS
	fmt.Println(cos)
	cmd := exec.Command("open", "http://localhost:9999")
	cmd.Run()

}
