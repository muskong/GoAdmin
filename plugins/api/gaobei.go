package main

import (
	"log"
	"time"

	"github.com/muskong/GoPkg/zaplog"
)

func init() {
	log.Println("插件init GaoBei")
}

type (
	_gaoBei struct {
		Today string
	}
	pluginInfo struct {
		Key   string
		Title string
		Value string
	}
)

func (a _gaoBei) Info(dest *[]any) error {
	zaplog.Sugar.Info(a.Today)

	*dest = append(*dest, pluginInfo{Key: "Title", Title: "接口名称", Value: "高贝"})
	*dest = append(*dest, pluginInfo{Key: "Class", Title: "包名称", Value: "GaoBei"})
	*dest = append(*dest, pluginInfo{Key: "PayPid", Title: "商户号", Value: ""})
	*dest = append(*dest, pluginInfo{Key: "PayKey", Title: "商户密钥", Value: ""})
	*dest = append(*dest, pluginInfo{Key: "PayName", Title: "支付宝收款姓名", Value: ""})
	*dest = append(*dest, pluginInfo{Key: "PayAccount", Title: "支付宝收款账户", Value: ""})
	*dest = append(*dest, pluginInfo{Key: "SendUrl", Title: "提交地址", Value: ""})
	*dest = append(*dest, pluginInfo{Key: "SearchUrl", Title: "查询地址", Value: ""})

	return nil
}

func (a _gaoBei) Send(request map[string]any, respond *map[string]any) error {
	*respond = request
	log.Println("GaoBei Send test", request, respond)
	return nil
}
func (a _gaoBei) Search(request map[string]any, respond *map[string]any) error {
	*respond = request
	log.Println("GaoBei Search test", request, respond)
	return nil
}
func (a _gaoBei) Notify(request map[string]any, respond *map[string]any) error {
	*respond = request
	log.Println("GaoBei Notify test", request, respond)
	return nil
}

var GaoBei = _gaoBei{
	Today: time.Now().Format("2006-01-02 15:04:05"),
}
