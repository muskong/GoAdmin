package logic

import (
	"plugin"
	"strings"

	"github.com/muskong/GoPkg/zaplog"
)

type _plugin struct {
}

var Plugin = &_plugin{}

func (p *_plugin) New(path string, config map[string]string) Api {
	api := p.Load(path)
	err := api.SetConfig(config)
	if err != nil {
		zaplog.Sugar.Error("plugin set config error : ", err)
		return nil
	}

	return api
}

func (p *_plugin) Load(path string) Api {
	plug, err := plugin.Open(path)
	if err != nil {
		zaplog.Sugar.Error("open fail : ", err)
	}
	zaplog.Sugar.Infof(" plug %+v", plug)

	_, lookupName, _ := strings.Cut(path, "_")
	lookupName = strings.TrimSuffix(lookupName, ".so")
	lookup, err := plug.Lookup(lookupName)
	if err != nil {
		zaplog.Sugar.Error("plugin name error : ", err)
		return nil
	}
	zaplog.Sugar.Infof(" plug.Lookup %#+v", lookup)
	api, ok := lookup.(Api)
	if !ok {
		zaplog.Sugar.Error("plugin interface fail")
		return nil
	}

	return api
}

func PluginGaoBei(path string) Api {
	plug, err := plugin.Open(path)
	if err != nil {
		zaplog.Sugar.Error("open fail : ", err)
	}

	lookup, err := plug.Lookup("GaoBei")
	if err != nil {
		zaplog.Sugar.Error("plugin name error : ", err)
		return nil
	}
	zaplog.Sugar.Infof(" plug.Lookup %+v", lookup)
	zaplog.Sugar.Infof(" plug.Lookup %#v", lookup)
	api, ok := lookup.(Api)
	if !ok {
		zaplog.Sugar.Error("plugin interface fail")
		return nil
	}

	return api
}

func PluginTest(path string) {
	p, err := plugin.Open(path)
	if err != nil {
		zaplog.Sugar.Error("open fail : ", err)
	}

	// _, lookupName, _ := strings.Cut(path, "_")
	// lookupName = strings.TrimSuffix(lookupName, ".so")
	// lookup, err := plug.Lookup(lookupName)

	v, err := p.Lookup("V")
	zaplog.Sugar.Infof(" plug.Lookup %+v", v)
	zaplog.Sugar.Infof(" plug.Lookup %#v", v)
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	zaplog.Sugar.Infof(" plug.Lookup %+v", f)
	zaplog.Sugar.Infof(" plug.Lookup %#v", f)
	if err != nil {
		panic(err)
	}
	*v.(*int) = 7
	f.(func())() // prints "Hello, number 7"
}
