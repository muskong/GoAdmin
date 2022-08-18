package external

type ShouKaYun struct {
	Api
	PayPid     string
	PayAccount string
}

func NewShouKaYun() *ShouKaYun {
	return &ShouKaYun{}
}

func (a *ShouKaYun) Info() ApiInfo {
	param := []map[string]string{}

	param = append(param, map[string]string{"name": "商户号", "field": "pay_pid"})
	param = append(param, map[string]string{"name": "商户密钥", "field": "pay_key"})
	param = append(param, map[string]string{"name": "支付宝收款姓名", "field": "pay_name"})
	param = append(param, map[string]string{"name": "支付宝收款账户", "field": "pay_account"})
	param = append(param, map[string]string{"name": "提交地址", "field": "sendurl"})
	param = append(param, map[string]string{"name": "查询地址", "field": "searchurl"})

	return ApiInfo{
		Title:     "收卡云",                        // 中文接口名称
		Name:      "ShouKaYun",                  // struct 名称
		Version:   "1.0.0",                      // 版本号
		Url:       "https://www.shoukayun.com/", // 提交数据的接口地址
		Parameter: param,                        // 接口参数
	}
}

func (a *ShouKaYun) encrypt(data string) string {
	return data
}
func (a *ShouKaYun) sign() {
	a.Data.Set("sign", "")
}
func (a *ShouKaYun) Send(data map[string]string) error {

	a.Data.Add("account", a.PayPid)
	a.Data.Add("order_no", data["order_number"])
	a.Data.Add("product_no", data["product_no"])
	a.Data.Add("card_no", a.encrypt(data["card_no"]))
	a.Data.Add("card_password", a.encrypt(data["card_password"]))
	a.Data.Add("notify_url", data["notify_url"])

	a.sign()

	var result map[string]any
	err := a.post(a.SendPath, &result)

	return err
}
