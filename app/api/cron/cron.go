package cron

import (
    "github.com/gogf/gf/container/gmap"
    "github.com/gogf/gf/net/ghttp"
    "github.com/gogf/gf/os/gcron"
    "github.com/gogf/gf/os/gtime"
    "github.com/gogf/gf/text/gstr"
    "github.com/gogf/gf/util/gconv"
    "kingProject/library/response"
    "kingProject/utils"
)

// 一次性定时任务对象
type C struct{}

// 全局变量
var (
    allCronContent = gmap.New()
)

func (c *C) AddCron(r *ghttp.Request) {
    var (
        data *AddCron
    )
    if err := r.Parse(&data); err != nil {
        response.JsonExit(r, 1, err.Error())
    }
    if cron := gcron.Search(data.CronName); cron != nil {
        response.JsonExit(r, 2, "任务名已存在！")
    }
    targetTime := data.TargetTime
    t, err := gtime.StrToTime(targetTime)
    if err != nil {
        response.JsonExit(r, 3, "目标时间格式错误！")
    }
    currentTime := gtime.Now()
    second := t.Sub(currentTime).Seconds()
    if second < 0 {
        response.JsonExit(r, 4, "目标时间晚于当前时间！")
    }
    // 配置规则：隔x秒，运行一次
    pattern := gstr.Join([]string{"@every ", gconv.String(gconv.Int(second)), "s"}, "")
    // set内容
    allCronContent.Set(data.CronName, data)
    gcron.AddOnce(pattern, func() {
        utils.SendWechat(data.Content)
        dingdingContent := `{"msgtype":"text","text":{"content":"来自云空\n` + data.Content + ` "}}`
        utils.SendDingDing(dingdingContent)
    }, data.CronName)
    response.JsonExit(r, 0, "ok", allCronContent.Get(data.CronName))
}

func (c *C) DeleteCron(r *ghttp.Request) {
    var (
        data *DeleteOrSearchCron
    )
    if err := r.Parse(&data); err != nil {
        response.JsonExit(r, 1, err.Error())
    }
    if cron := gcron.Search(data.CronName); cron == nil {
        response.JsonExit(r, 2, "任务不存在！")
    } else {
        gcron.Remove(data.CronName)
        response.JsonExit(r, 0, "ok")
    }
}

func (c *C) GetOneCron(r *ghttp.Request) {
    var (
        data *DeleteOrSearchCron
    )
    if err := r.Parse(&data); err != nil {
        response.JsonExit(r, 1, err.Error())
    }
    if cron := gcron.Search(data.CronName); cron == nil {
        response.JsonExit(r, 2, "任务不存在！")
    } else {
        response.JsonExit(r, 0, "ok", allCronContent.Get(data.CronName))
    }
}

func (c *C) GetAllCron(r *ghttp.Request) {
    response.JsonExit(r, 0, "ok", allCronContent.Values())
}
