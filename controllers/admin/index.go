package admin

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

//首页
func (c *IndexController) Get() {

	c.TplName = "admin/index.html"
}
