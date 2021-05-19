package alcor

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"kingProject/library/response"
	"kingProject/utils"
)

type C struct{}

func (c *C) GetTlmPrice(r *ghttp.Request) {
	result := gmap.New()
	HttpGetTlmPrice(result)
	response.JsonExit(r, 0, "ok", result)
}

func HttpGetTlmPrice(result *gmap.Map) {
	if j, err := g.Client().Get("https://wax.alcor.exchange/api/markets"); err != nil {
		panic(err)
	} else {
		defer j.Close()
		allString := j.ReadAllString()
		g.Log().Info(allString)
		x := gjson.New(allString)
		y := x.Array()
		for _, i := range y {
			z := gjson.New(i)
			if z.Get("id") != nil && z.GetInt("id") == 26 {
				result.Set("id", z.GetInt("id"))
				result.Set("lastPrice", z.GetFloat64("last_price"))
				result.Set("change24", z.GetFloat64("change24"))
				result.Set("name", z.GetString("quote_token.symbol.name"))
			}
		}
	}
}

func MonitorPrice() {
	gcron.Add("@every 1h", func() {
		result := gmap.New()
		HttpGetTlmPrice(result)
		lastPrice := result.Get("lastPrice")
		monitorTlmPriceThreshold := g.Cfg().GetString("custom.monitorTlmPriceThreshold")
		if lastPrice != nil && gconv.Float64(lastPrice) > gconv.Float64(monitorTlmPriceThreshold) {
			wechatContent := gstr.Join([]string{"tlm当前价格:", gconv.String(lastPrice), ",超过了:", monitorTlmPriceThreshold, "当前时间：", gtime.Now().String()}, "")
			utils.SendWechat(wechatContent)
			dingdingContent := `{"msgtype":"text","text":{"content":"来自云空\n` + wechatContent + ` "}}`
			utils.SendDingDing(dingdingContent)
			g.Log().Info(wechatContent)
		}
	})

}
