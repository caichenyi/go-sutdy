package router

import (
    "github.com/astaxie/beego"
    "SecKill/SecProxy/controller"
)

func init() {
    beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")
    beego.Router("/secinfo", &SkillController{}, "*;SecInfo")
}