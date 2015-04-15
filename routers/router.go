package routers

import (
	"jianlibang/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/resume", &controllers.ResumeController{})
}
