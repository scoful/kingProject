package coingecko

import (
	"github.com/gogf/gf/container/garray"
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

func (c *C) GetPrice(r *ghttp.Request) {
	var (
		data *GetPrice
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	result := garray.New()
	HttpGetPrice(result, data)
	response.JsonExit(r, 0, "ok", result)
}

func HttpGetPrice(result *garray.Array, data *GetPrice) {
	url := gstr.Join([]string{"https://api.coingecko.com/api/v3/simple/price?ids=", data.CoinIds, "&vs_currencies=", data.VsCurrencie}, "")
	if j, err := g.Client().Get(url); err != nil {
		panic(err)
	} else {
		defer j.Close()
		allString := j.ReadAllString()
		g.Log().Info(allString)
		x := gjson.New(allString)
		strs := gstr.Split(data.CoinIds, ",")
		for _, s := range strs {
			coinMap := gmap.New()
			coinMap.Set("coinId", s)
			coinMap.Set("vsCurrencie", data.VsCurrencie)
			coinMap.Set("price", x.GetFloat64(gstr.Join([]string{s, data.VsCurrencie}, ".")))
			result.Append(coinMap)
		}
	}
}

// 暂时只监控idena价格，后续要接入mysql，再搞多个币价监控
func MonitorPrice() {
	gcron.Add("@every 1h", func() {
		if j, err := g.Client().Get("https://api.coingecko.com/api/v3/simple/price?ids=idena&vs_currencies=cny"); err != nil {
			panic(err)
		} else {
			defer j.Close()
			allString := j.ReadAllString()
			g.Log().Info(allString)
			x := gjson.New(allString)
			lastPrice := x.Get("idena.cny")
			if lastPrice != nil && gconv.Float64(lastPrice) > 1.5 {
				wechatContent := gstr.Join([]string{"idena当前价格:", gconv.String(lastPrice), ",超过了:￥1.5", "当前时间：", gtime.Now().String()}, "")
				utils.SendWechat(wechatContent)
				dingdingContent := `{"msgtype":"text","text":{"content":"来自云空\n` + wechatContent + ` "}}`
				utils.SendDingDing(dingdingContent)
				g.Log().Info(wechatContent)
			}
		}
	})

}
