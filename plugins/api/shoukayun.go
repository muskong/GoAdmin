package main

import (
	"errors"
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
		Key   string
		Title string
		Value string
	}
)

func (a shouKaYun) Info(dest *[]any) error {
	zaplog.Sugar.Info(a.Today)

	*dest = append(*dest, pluginInfo{Key: "Title", Title: "接口名称", Value: "收卡云"})
	*dest = append(*dest, pluginInfo{Key: "Class", Title: "包名称", Value: "ShouKaYun"})
	*dest = append(*dest, pluginInfo{Key: "PayPid", Title: "商户号", Value: ""})
	*dest = append(*dest, pluginInfo{Key: "PayKey", Title: "商户密钥", Value: ""})
	*dest = append(*dest, pluginInfo{Key: "PayName", Title: "支付宝收款姓名", Value: ""})
	*dest = append(*dest, pluginInfo{Key: "PayAccount", Title: "支付宝收款账户", Value: ""})
	*dest = append(*dest, pluginInfo{Key: "SendUrl", Title: "提交地址", Value: ""})
	*dest = append(*dest, pluginInfo{Key: "SearchUrl", Title: "查询地址", Value: ""})

	return nil
}

func (a shouKaYun) encrypt(data string) string {
	return data
}
func (a shouKaYun) sign() {
	a.Data.Set("sign", "")
}

func (a shouKaYun) Send(request map[string]string, respond *map[string]any) error {

	a.Data.Add("account", a.PayPid)
	a.Data.Add("order_no", request["order_number"])
	a.Data.Add("product_no", request["product_no"])
	a.Data.Add("card_no", a.encrypt(request["card_no"]))
	a.Data.Add("card_password", a.encrypt(request["card_password"]))
	a.Data.Add("notify_url", request["notify_url"])

	a.sign()

	// err := post(a.SendPath, respond)

	return nil
}
func (p shouKaYun) Search(request map[string]any, respond *map[string]any) error {
	*respond = request
	return errors.New("shouKaYun Search test")
}
func (p shouKaYun) Notify(request map[string]any, respond *map[string]any) error {
	*respond = request
	return errors.New("shouKaYun Notify test")
}

var ShouKaYun = shouKaYun{
	Today: time.Now().Format("2006-01-02 15:04:05"),
}
