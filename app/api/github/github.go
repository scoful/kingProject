package github

import (
    "github.com/gogf/gf/encoding/gjson"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"
    "github.com/gogf/gf/os/gtime"
    "github.com/gogf/gf/text/gstr"
    "kingProject/library/response"
    "kingProject/utils"
)

// github对象
type C struct{}

func (c *C) HandleGithubWebhook(r *ghttp.Request) {
    if j, err := gjson.DecodeToJson(r.GetBodyString()); err != nil {
        panic(err)
    } else {
        action := j.GetString("action")
        // ping判断
        zen := j.GetJson("zen")
        if !zen.IsNil() {
            action = "ping"
        }
        // fork判断
        forkee := j.GetJson("forkee")
        if !forkee.IsNil() {
            action = "fork"
        }
        // push判断
        pusher := j.GetJson("pusher")
        if !pusher.IsNil() {
            action = "push"
        }
        // 通用
        sender := j.GetJson("sender")
        login := sender.GetString("login")
        html_url := sender.GetString("html_url")
        repository := j.GetJson("repository")
        repositoryName := repository.GetString("name")
        switch {
        case action == "deleted":
            wechatContent := gstr.Join([]string{login, " 取消关注了 ", repositoryName, "\n地址是：", html_url, "\n", gtime.Now().String()}, "")
            utils.SendWechat(wechatContent)
            dingdingContent := `{"msgtype":"text","text":{"content":"来自云空\n` + wechatContent + ` "}}`
            utils.SendDingDing(dingdingContent)
            g.Log().Info(wechatContent)
        case action == "started":
            wechatContent := gstr.Join([]string{login, " star了 ", repositoryName, "\n地址是：", html_url, "\n", gtime.Now().String()}, "")
            utils.SendWechat(wechatContent)
            dingdingContent := `{"msgtype":"text","text":{"content":"来自云空\n` + wechatContent + ` "}}`
            utils.SendDingDing(dingdingContent)
            g.Log().Info(wechatContent)
        case action == "fork":
            wechatContent := gstr.Join([]string{login, " fork了 ", repositoryName, "\n地址是：", html_url, "\n", gtime.Now().String()}, "")
            utils.SendWechat(wechatContent)
            dingdingContent := `{"msgtype":"text","text":{"content":"来自云空\n` + wechatContent + ` "}}`
            utils.SendDingDing(dingdingContent)
            g.Log().Info(wechatContent)
        case action == "push":
            wechatContent := gstr.Join([]string{login, " push了 ", repositoryName, "\n地址是：", html_url, "\n", gtime.Now().String()}, "")
            utils.SendWechat(wechatContent)
            dingdingContent := `{"msgtype":"text","text":{"content":"来自云空\n` + wechatContent + ` "}}`
            utils.SendDingDing(dingdingContent)
            g.Log().Info(wechatContent)
        case action == "ping":
            wechatContent := gstr.Join([]string{repositoryName, " 成功接入了事件通知\n操作者：", login, "\n", gtime.Now().String()}, "")
            utils.SendWechat(wechatContent)
            dingdingContent := `{"msgtype":"text","text":{"content":"来自云空\n` + wechatContent + ` "}}`
            utils.SendDingDing(dingdingContent)
            g.Log().Info(wechatContent)
        default:
            wechatContent := gstr.Join([]string{login, " 操作了 ", repositoryName, "，但不知道是啥操作\n地址是：", html_url, "\n", gtime.Now().String()}, "")
            utils.SendWechat(wechatContent)
            dingdingContent := `{"msgtype":"text","text":{"content":"来自云空\n` + wechatContent + ` "}}`
            utils.SendDingDing(dingdingContent)
            g.Log().Info(wechatContent)
            g.Log().Info(r.GetBodyString())
        }
    }
    response.JsonExit(r, 0, "ok")
}
