package logic

import (
	"log"
	"plugin"
	"strings"
)

func Plugin(path string) Api {
	plug, err := plugin.Open(path)
	if err != nil {
		log.Println("open fail : ", err)
	}

	_, lookupName, _ := strings.Cut(path, "_")
	lookupName = strings.TrimSuffix(lookupName, ".so")
	lookup, err := plug.Lookup(lookupName)
	if err != nil {
		log.Println("plugin name error : ", err)
		return nil
	}
	api, ok := lookup.(Api)
	if !ok {
		log.Println("plugin interface fail")
		return nil
	}

	return api
}
