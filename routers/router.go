package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"xhcms/controllers"
	"xhcms/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &admin.LoginController{}, "get:GetLogin;post:PostLogin")
	beego.Router("/admin/index", &admin.IndexController{})
	beego.Handler("/captcha/*.png", captcha.Server(150, 36)) //注册验证码服务，验证码图片高度150x36
}
