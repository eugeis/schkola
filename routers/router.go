package routers

import (
	"github.com/eugeis/schkola/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
