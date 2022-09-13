package main

import (
	"log"
	"time"
)

type BadNastyDocter string

func (g BadNastyDocter) HealthCheck() error {
	log.Println("now is", g)

	return nil
}

//go build -buildmode=plugin -o plugins/pay_Alipay.so plugins/pay/alipay.go

var Alipay = BadNastyDocter(time.Now().Format("2006-01-02 15:04:05"))
