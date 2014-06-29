package routers

import (
	"github.com/astaxie/beego"
	//"github.com/liangdian/waroengbali.com/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.SetStaticPath("/static", "static")
}
