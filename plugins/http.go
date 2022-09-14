package plugin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/muskong/GoPkg/zaplog"
)

type (
	PluginApi struct {
		ContentType string // json|form
		Data        url.Values
		Url         string // 接口域名或IP
		ApiID       string // 商户号
		ApiKey      string // 商户密钥
	}
)

func (a PluginApi) post(path string, result *map[string]any) (err error) {
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

	err = json.Unmarshal(rspBody, *result)
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	return
}
