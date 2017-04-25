package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"xhcms/models/sqlite"
)

type LoginController struct {
	beego.Controller
}

//登录首页
func (c *LoginController) GetLogin() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "admin/login.html"
}

//提交
func (c *LoginController) PostLogin() {
	username := c.GetString("username")
	password := c.GetString("password")
	valid := validation.Validation{}
	valid.Required(username, "请输入用户名")
	valid.MinSize(username, 5, "用户名最少5位")
	valid.MaxSize(username, 15, "用户名最多15位")
	valid.Required(password, "请输入密码")
	valid.MinSize(password, 8, "密码最少8位")
	valid.MaxSize(password, 15, "密码最多15位")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["josn"] = map[string]interface{}{"status": 4, "info": err.Key}
			c.ServeJSON()
			return
		}
	} else {
		has, err := sqlite.LoginUser(username, password)
		if err != nil {
			c.Data["json"] = map[string]interface{}{"status": 4, "info": "发生未知错误"}
			c.ServeJSON()
			return
		} else {
			if has == true {
				c.Data["json"] = map[string]interface{}{"status": 1, "info": "登录成功！3秒后为你跳转！"}
				c.ServeJSON()
				return
			} else {
				c.Data["json"] = map[string]interface{}{"status": 4, "info": "用户名或密码错误！"}
				c.ServeJSON()
				return
			}
		}
	}
}
