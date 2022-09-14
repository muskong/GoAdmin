package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/forgoer/openssl"
	"github.com/muskong/GoPkg/plugin"
)

func init() {
	log.Println("插件init ShouKaYun")
}

type (
	shouKaYun struct {
		Data url.Values
		// PayPid     string
		// PayKey     string
		// PayAccount string
		Http       plugin.PluginApi
		Title      string
		Class      string
		PayPid     string
		PayKey     string
		PayName    string
		PayAccount string
		SendUrl    string
		SearchUrl  string
	}
	pluginInfo struct {
		Name     string
		Label    string
		Value    string
		Disabled bool
	}
)

func (a *shouKaYun) Info(dest *[]any) error {

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

func (a *shouKaYun) encrypt(data string) string {
	h := md5.New()
	crypto := h.Sum([]byte(a.PayKey))
	iv := hex.EncodeToString(crypto)
	iv = iv[8:24]

	dst, _ := openssl.AesCBCEncrypt([]byte(data), []byte(a.PayKey), []byte(iv), openssl.PKCS7_PADDING)
	return base64.StdEncoding.EncodeToString(dst)
}
func (a *shouKaYun) sign() {
	ksort := []string{}
	for k, _ := range a.Data {
		ksort = append(ksort, k)
	}

	sort.Strings(ksort)

	var str string
	for _, v := range ksort {
		if v == "sign" || v == "account" {
			continue
		}
		data := a.Data.Get(v)
		if len(data) == 0 {
			continue
		}
		str += fmt.Sprintf("%s=%s&", v, data)
	}
	str += fmt.Sprintf("key=%s", a.PayKey)

	h := md5.New()
	crypto := h.Sum([]byte(str))
	sign := strings.ToUpper(hex.EncodeToString(crypto))

	a.Data.Set("sign", sign)
}

func (a *shouKaYun) SetConfig(conf map[string]string) {
	a.Title = conf["Title"]
	a.Class = conf["Class"]
	a.PayPid = conf["PayPid"]
	a.PayKey = conf["PayKey"]
	a.PayName = conf["PayName"]
	a.PayAccount = conf["PayAccount"]
	a.SendUrl = conf["SendUrl"]
	a.SearchUrl = conf["SearchUrl"]
}

func (a *shouKaYun) Send(request map[string]string, respond any) error {
	resp := struct {
		Message string
		Data    any
		Time    time.Time
	}{
		Time: time.Now(),
	}

	a.Data.Add("account", a.PayPid)
	a.Data.Add("order_no", request["order_number"])
	a.Data.Add("product_no", request["product_no"])
	a.Data.Add("card_no", a.encrypt(request["card_no"]))
	a.Data.Add("card_password", a.encrypt(request["card_password"]))
	a.Data.Add("notify_url", request["notify_url"])

	a.sign()
	var result map[string]any
	err := a.Http.Post(a.SendUrl, &result)
	if err == nil {
		resp.Message = err.Error()
	}

	// if

	respond = &resp
	return err
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
var ShouKaYun = &shouKaYun{
	// Today: time.Now().Format("2006-01-02 15:04:05"),
	Http: plugin.PluginApi{},
}
