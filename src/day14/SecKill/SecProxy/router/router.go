package router

import (
    "day14/SecKill/SecProxy/controller"
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego"
)

func init() {
    logs.Debug("enter route init")
    beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")
    beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo")
}