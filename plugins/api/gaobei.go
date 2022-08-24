package main

import (
	"log"
)

func init() {
	log.Println("插件init GaoBei")
}

type _gaoBei struct{}

func (_gaoBei) Info() map[string]string {
	info := map[string]string{}
	info["Title"] = "高贝"     // 中文接口名称
	info["Class"] = "GaoBei" // 中文接口名称
	info["payPid"] = ""      // 商户号
	info["payKey"] = ""      // 商户密钥
	info["payName"] = ""     // 支付宝收款姓名
	info["payAccount"] = ""  // 支付宝收款账户
	info["sendurl"] = ""     // 提交地址
	info["searchurl"] = ""   // 查询地址

	return info
}

func (_gaoBei) Send(request map[string]any, respond *map[string]any) error {
	*respond = request
	log.Println("GaoBei Send test", request, respond)
	return nil
}
func (_gaoBei) Search() error {
	log.Println("GaoBei Search test")
	return nil
}
func (_gaoBei) Notify() error {
	log.Println("GaoBei Notify test")
	return nil
}

var GaoBei = _gaoBei{}
