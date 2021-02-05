package utils

import (
    "github.com/gogf/gf/encoding/gurl"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/text/gstr"
)

func SendWechat(content string) {
    wechatUrl := g.Cfg().GetString("custom.wechatUrl")
    url := gstr.Join([]string{wechatUrl, gurl.Encode(content)}, "")
    if r, err := g.Client().Get(url); err != nil {
        panic(err)
    } else {
        defer r.Close()
        g.Log().Info(r.ReadAllString())
    }
}

func SendDingDing(content string) {
    dingdingUrl := g.Cfg().GetString("custom.dingdingUrl")
    c := g.Client()
    c.SetHeader("Content-Type", "application/json")
    if r, err := c.Post(
        dingdingUrl,
        content,
    ); err != nil {
        panic(err)
    } else {
        defer r.Close()
        g.Log().Info(r.ReadAllString())
    }
}
