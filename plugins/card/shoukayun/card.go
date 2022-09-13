package main

import (
	"log"
	"net/url"
	"time"

	"github.com/muskong/GoPkg/zaplog"
)

func init() {
	log.Println("插件init ShouKaYun")
}

type (
	shouKaYun struct {
		Today      string
		Data       url.Values
		SendPath   string // 提交数据的接口地址
		SearchPath string // 查询数据地址
		PayPid     string
		PayAccount string
	}
	pluginInfo struct {
		Name     string
		Label    string
		Value    string
		Disabled bool
	}
)

func (a shouKaYun) Info(dest *[]any) error {
	zaplog.Sugar.Info(a.Today)

	*dest = append(*dest, &pluginInfo{Name: "Title", Label: "接口名称", Value: "收卡云", Disabled: true})
	*dest = append(*dest, &pluginInfo{Name: "Class", Label: "包名称", Value: "ShouKaYun", Disabled: true})
	*dest = append(*dest, &pluginInfo{Name: "PayPid", Label: "商户号", Value: "", Disabled: false})
	*dest = append(*dest, &pluginInfo{Name: "PayKey", Label: "商户密钥", Value: "", Disabled: false})
	*dest = append(*dest, &pluginInfo{Name: "PayName", Label: "支付宝收款姓名", Value: "", Disabled: false})
	*dest = append(*dest, &pluginInfo{Name: "PayAccount", Label: "支付宝收款账户", Value: "", Disabled: false})
	*dest = append(*dest, &pluginInfo{Name: "SendUrl", Label: "提交地址", Value: "", Disabled: false})
	*dest = append(*dest, &pluginInfo{Name: "SearchUrl", Label: "查询地址", Value: "", Disabled: false})

	return nil
}

func (a shouKaYun) encrypt(data string) string {
	return data
}
func (a shouKaYun) sign() {
	a.Data.Set("sign", "")
}

func (a shouKaYun) Send(request map[string]any, respond *map[string]any) error {

	// a.Data.Add("account", a.PayPid)
	// a.Data.Add("order_no", request["order_number"])
	// a.Data.Add("product_no", request["product_no"])
	// a.Data.Add("card_no", a.encrypt(request["card_no"]))
	// a.Data.Add("card_password", a.encrypt(request["card_password"]))
	// a.Data.Add("notify_url", request["notify_url"])

	// a.sign()

	// err := post(a.SendPath, respond)

	return nil
}
func (p shouKaYun) Search(request map[string]any, respond *map[string]any) error {
	*respond = request
	log.Println("GaoBei Search test", request, respond)
	return nil
}
func (p shouKaYun) Notify(request map[string]any, respond *map[string]any) error {
	*respond = request
	log.Println("GaoBei Notify test", request, respond)
	return nil
}

// go build -buildmode=plugin -o plugins/card_ShouKaYun.so plugins/card/shoukayun/card.go
var ShouKaYun = shouKaYun{
	Today: time.Now().Format("2006-01-02 15:04:05"),
}
