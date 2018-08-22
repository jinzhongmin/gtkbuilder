package main

import (
	"os"

	"github.com/jinzhongmin/gtkbuilder"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)

	//go-gtk builder
	builder := gtk.NewBuilder()
	builder.AddFromFile("main.ui")

	//gtkbuilder builder
	mbuilder := gtkbuilder.NewBuilder(builder)
	mainWindow := mbuilder.GetWindow("window1")
	mainWindow.SetResizable(false)
	mainWindow.Connect("destroy", gtk.MainQuit)

	//事件响应
	mbuilder.Connect("but-clicked", func() {
		name := mbuilder.GetEntry("entry1").GetText()
		if len(name) != 0 {
			mbuilder.GetLabel("label1").SetText("你好 " + name)
		} else {
			mbuilder.GetLabel("label1").SetText("请输入你的名字")
			mbuilder.GetLabel("entry1").GrabFocus()
		}

	})

	//开启事件监听
	mbuilder.Listening()

	mainWindow.ShowAll()
	gtk.Main()
}
