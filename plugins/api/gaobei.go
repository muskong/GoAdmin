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
		Name     string
		Label    string
		Value    string
		Disabled bool
	}
)

func (a _gaoBei) Info(dest *[]any) error {
	zaplog.Sugar.Info(a.Today)

	*dest = append(*dest, &pluginInfo{Name: "Title", Label: "接口名称", Value: "高贝", Disabled: true})
	*dest = append(*dest, &pluginInfo{Name: "Class", Label: "包名称", Value: "GaoBei", Disabled: true})
	*dest = append(*dest, &pluginInfo{Name: "PayPid", Label: "商户号", Value: "", Disabled: false})
	*dest = append(*dest, &pluginInfo{Name: "PayKey", Label: "商户密钥", Value: "", Disabled: false})
	*dest = append(*dest, &pluginInfo{Name: "PayName", Label: "支付宝收款姓名", Value: "", Disabled: false})
	*dest = append(*dest, &pluginInfo{Name: "PayAccount", Label: "支付宝收款账户", Value: "", Disabled: false})
	*dest = append(*dest, &pluginInfo{Name: "SendUrl", Label: "提交地址", Value: "", Disabled: false})
	*dest = append(*dest, &pluginInfo{Name: "SearchUrl", Label: "查询地址", Value: "", Disabled: false})

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
