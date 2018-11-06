package main

import "gitee.com/johng/gf/g/net/ghttp"

func main() {
	s := ghttp.GetServer()
	s.BindHandler("/", func(r *ghttp.Request){
		r.Response.WriteString("Hello World!")
	})
	s.Run()
}

