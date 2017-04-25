package routers

import (
	"github.com/astaxie/beego"
	"xhcms/controllers"
	"xhcms/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &admin.LoginController{}, "get:GetLogin;post:PostLogin")
	beego.Router("/admin/index", &admin.IndexController{})
}
