package router

import (
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"
    "kingProject/app/api/cron"
    "kingProject/app/api/github"
    "kingProject/app/service/middleware"
)

func init() {
    s := g.Server()

    s.Group("/", func(group *ghttp.RouterGroup) {
        group.Middleware(middleware.CORS)
        ctlCron := new(cron.C)
        ctlGithub := new(github.C)

        // 一次性定时提醒
        group.GET("/getAllCron", ctlCron, "GetAllCron")
        group.POST("/addCron", ctlCron, "AddCron")
        group.DELETE("/deleteCron", ctlCron, "DeleteCron")
        group.GET("/getOneCron", ctlCron, "GetOneCron")

        // githubwebhook
        group.POST("/handleGithubWebhook", ctlGithub, "HandleGithubWebhook")
    })

}
