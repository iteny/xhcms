package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/dchest/captcha"
	"xhcms/models"
	"xhcms/models/sqlite"
)

type LoginController struct {
	beego.Controller
}

//登录首页
func (c *LoginController) GetLogin() {
	captchaId := captcha.NewLen(4) //验证码长度
	c.Data["CaptchaId"] = captchaId
	c.TplName = "admin/login.html"
}

//提交
func (c *LoginController) PostLogin() {
	id, value := c.GetString("captcha_id"), c.GetString("captcha")
	b := captcha.VerifyString(id, value)
	if b == true {
		username, password := c.GetString("username"), c.GetString("password")
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
			row, err := sqlite.UserDb.LoginUser(username, password)
			fmt.Println(row)
			if err != nil {
				c.Data["json"] = map[string]interface{}{"status": 4, "info": "发生未知错误"}
				c.ServeJSON()
				return
			} else {
				if row.Id != 0 {
					ip := c.Ctx.Input.IP()
					ipinfo := models.TaobaoIP(ip)
					if row.Status == 0 {
						c.Data["json"] = map[string]interface{}{"status": 4, "info": "帐号已被禁用！"}
						c.ServeJSON()
						return
					} else if row.Status == 1 {
						sqlite.LoginLogDb.RecordLogin(row.Username, ip, 1, "登录成功", ipinfo.Data.Area, ipinfo.Data.Country, c.Ctx.Input.UserAgent())
						c.SetSession("userid", row.Id)
						c.SetSession("username", row.Username)
						c.SetSession("status", row.Status)
						c.Data["json"] = map[string]interface{}{"status": 1, "info": "登录成功！3秒后为你跳转！"}
						c.ServeJSON()
						return
					} else {
						c.Data["json"] = map[string]interface{}{"status": 4, "info": "发生未知错误"}
						c.ServeJSON()
						return
					}

				} else {
					c.Data["json"] = map[string]interface{}{"status": 4, "info": "用户名或密码错误！"}
					c.ServeJSON()
					return
				}
			}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"status": 4, "info": "验证码错误！"}
		c.ServeJSON()
		return
	}

}
