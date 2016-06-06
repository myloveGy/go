package controllers

import (
	"github.com/astaxie/beego"
)

// 用户数据
type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
}

// 请求返回数据
type ArrError struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// 继承基础控制器
type BaseController struct {
	beego.Controller
	U User
	E ArrError
}

// 验证用户是否已经登录
func (this *BaseController) isLogin(str string) bool {
	user := this.GetSession(str)
	if user != nil {
		this.U = user.(User)
		if this.U.Id > 0 {
			return true
		}
	}

	return false
}

// 返回json数据
func (this *BaseController) AjaxReturn() {
	this.Data["json"] = &this.E
	this.ServeJSON()
}
