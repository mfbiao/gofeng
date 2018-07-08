package gofeng

import (
	"os"
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/glib"
)

func HandleSignal(ctx *glib.CallbackContext){
	fmt.Println("^_^ ^_^ ^_^")
	arg := ctx.Data() //获取用户传递的参数，是个空接口类型
	data, ok := arg.(string) //类型断言
	if ok {
		fmt.Println(data)
	}
}

func main() {
	//一。初始化 固定
	gtk.Init(&os.Args)
	//二。用户代码
	//1。创建窗口 2。设置属性，3。显示窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //带边框的顶层窗口
	win.SetTitle("go gtk")                    //设置标题
	win.SetSizeRequest(480, 320)              //设置大小
	//3.创建容器控件（固定布局，任意布局）
	layout := gtk.NewFixed()
	//4.布局添加到窗口上
	win.Add(layout)
	//5。创建按钮
	btn1 := gtk.NewButtonWithLabel("^_^")
	btn2 := gtk.NewButtonWithLabel("@_@")
	btn2.SetSizeRequest(40,40)
	//6.按钮添加到布局中
	layout.Put(btn1,20,20)
	layout.Put(btn2,100,100)
	//7.显示控件
	//如果多个控件，如果使用show  需要每个都show
	//win.Show()
	//layout.Show()
	//btn1.Show()
	//也可以ShowAll
	win.ShowAll()
	//8.信号处理
	//按键按下处罚的信号"checked"
    str := "are u ok?"
    //Connect()只会调用一次
    btn1.Connect("clicked", HandleSignal, str)
    //处理函数可以匿名函数
    //btn2.Connect("clicked", func() {
    //	fmt.Println("=========")
    //    fmt.Println(str) // 可以直接用变量
	//})
    //第二种方式
    btn2.Clicked(func() {
		fmt.Println("@_@ @_@ @_@")
		fmt.Println(str) // 可以直接用变量
	})
    //关闭按钮 触发 destroy
    win.Connect("destroy", func() {
		fmt.Println("退出")
    	gtk.MainQuit()
	})
	//三，主事件循环（固定）
	//a。让程序不结束，b,等待用户操作 （移动窗口，点击鼠标）
	gtk.Main()
}


