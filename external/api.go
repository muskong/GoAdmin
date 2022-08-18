package external

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/muskong/GoPkg/zaplog"
)

type (
	ApiInterface interface {
		Send() error   // 提交数据
		Search()       // 查询数据
		Notify()       // 回调处理
		Config()       // 配置数据
		Info() ApiInfo // 接口信息
	}

	Api struct {
		ContentType string // json|form
		Data        url.Values
		Url         string // 接口域名或IP
		SendPath    string // 提交数据的接口地址
		SearchPath  string // 查询数据地址
		ApiID       string // 商户号
		ApiKey      string // 商户密钥
	}

	// 接口信息
	ApiInfo struct {
		Title     string              // 中文接口名称
		Name      string              // struct 名称
		Version   string              // 版本号
		Url       string              // 提交数据的接口地址
		Parameter []map[string]string // 接口参数
	}
)

func (a *Api) post(path string, result *map[string]any) (err error) {
	u, err := url.ParseRequestURI(a.Url)
	u.Path = path

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(a.Data.Encode()))
	if err != nil {
		zaplog.Sugar.Error(err)
		return
	}

	contentType := "application/x-www-form-urlencoded"
	if strings.ToUpper(a.ContentType) == "JSON" {
		contentType = "application/json"
	}
	req.Header.Add("Content-Type", contentType)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		zaplog.Sugar.Error(err)
		return
	}
	defer rsp.Body.Close()

	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		zaplog.Sugar.Error(err)
		return
	}

	err = json.Unmarshal(rspBody, &result)
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	return
}
