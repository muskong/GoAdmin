package main

import (
	"errors"
	"log"
)

func init() {
	log.Println("插件init ShouKaYun")
}

type shouKaYun struct {
	Api
	SendPath   string // 提交数据的接口地址
	SearchPath string // 查询数据地址
	PayPid     string
	PayAccount string
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

	err := a.post(a.SendPath, respond)

	return err
}
func (p shouKaYun) Search(request map[string]any, respond *map[string]any) error {
	*respond = request
	return errors.New("shouKaYun Search test")
}
func (p shouKaYun) Notify(request map[string]any, respond *map[string]any) error {
	*respond = request
	return errors.New("shouKaYun Notify test")
}

var ShouKaYun = shouKaYun{}
