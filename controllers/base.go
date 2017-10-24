package controllers

import (
	"blogapi/util"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
	options        map[string]string
	cache          *util.MyCache
}

func (this *BaseController) SendJSON(code int, data interface{}, msg string) {
	out := make(map[string]interface{})
	out["code"] = code
	out["data"] = data
	out["msg"] = msg
	this.Data["json"] = out
	this.ServeJSON()
}

func (this *BaseController) CheckError(err error) {
	if err != nil {
		this.SendJSON(400, "", err.Error())
	}
}
