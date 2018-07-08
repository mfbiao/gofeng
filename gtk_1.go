package gofeng

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//一。初始化 固定
	gtk.Init(&os.Args)
	//二。用户代码
	//1。创建窗口 2。设置属性，3。显示窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //带边框的顶层窗口
	win.SetTitle("go gtk")                    //设置标题
	win.SetSizeRequest(480, 320)              //设置大小
	win.Show()


	//三，主事件循环（固定）
	//a。让程序不结束，b,等待用户操作 （移动窗口，点击鼠标）
	gtk.Main()

}
