package main

import (
	//"fmt"
	"log"
	"net/http"
	"runtime"
	"os/exec"
	//"strings"
	//"os"
	//"html/template"
	"io"
	"time"
)
type Todo struct {
	Task string
	Done bool
}

const (
	PAOT  string = ":9999"
)
func main() {
	go openWindow()
	//http.Handle("/static/", http.FileServer(http.Dir("template")))
	//http.HandleFunc("/static", staticHandle)           //设置访问的路由
	http.HandleFunc("/", logPanics(index))           //设置访问的路由
	err := http.ListenAndServe(PAOT, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
// http静态文件
func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, "hello world!")
	case "POST":
		r.ParseForm()
		io.WriteString(w, r.Form["in"][1])
		io.WriteString(w, "\n")
		io.WriteString(w, r.FormValue("in"))
	}
	
	//static_dir := "./static";
	//if strings.HasPrefix(r.URL.Path,"/static"){
	//	file := static_dir + r.URL.Path[len("/static"):]
	//	f,err := os.Open(file)
	//	defer f.Close()
	//
	//	if err != nil && os.IsNotExist(err){
	//		file = static_dir + "/default.jpg"
	//	}
	//	http.ServeFile(w,r,file)
	//	return
	//}else {
	//	tmpl := template.Must(template.ParseFiles("./todos.html"))
	//	todos := []Todo{
	//		{"Learn Go", true},
	//		{"Read Go Web Examples", true},
	//		{"Create a web app in Go", false},
	//	}
	//	//println(html)
	//	tmpl.Execute(w, struct{ Todos []Todo }{todos})
	//}
}

//错误处理函数 获取异常
func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}

//打开浏览器
func openWindow(){
	time.Sleep(2 * time.Second) //延时执行
	cos := runtime.GOOS
	if cos == "windows" {
		cmd := exec.Command("cmd", "/C", "start http://localhost"+PAOT+"/index")
		cmd.Run()
	} else if cos == "mac" || cos == "darwin" {
		cmd := exec.Command("open", "http://localhost" + PAOT)
		cmd.Run()
	} else {
		cmd := exec.Command("xdg-open", "http://localhost" + PAOT)
		cmd.Run()
	}

}