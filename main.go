package main

import (
	"github.com/astaxie/beego"
	_ "github.com/liangdian/waroengbali.com/routers"
)

func main() {
	beego.DirectoryIndex = true
	beego.Run()
}
