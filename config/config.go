package config

import (
	"flag"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/nporsche/smartconfig/client"
)

var (
	configFile string
)
var GlobalConfig Config

/*
mode = "debug"

http_listen_address = ":8001"
grpc_listen_address = ":9095"

http_notify_host = "http://pay.chongdonglvxing.com/"

[wxpay]
api_address = "https://api.mch.weixin.qq.com/pay/unifiedorder"
api_id = ""
key = ""
notify_path = "/wxpay/cb"

[alipay]
api_address = "https://openapi.alipay.com/gateway.do"
api_id = ""
key = ""
notify_path = "/alipay/cb"
*/

type WxPay struct {
	APIAddress string `toml:"api_address"`
	APIID      string `toml:"api_id"`
	SellerID   string `toml:"seller_id"`
	Key        string `toml:"key"`
	NotifyPath string `toml:"notify_path"`
}

type AliPay struct {
	APIAddress string `toml:"api_address"`
	APIID      string `toml:"api_id"`
	PrivateKey string `toml:"private_key"`
	PublicKey  string `toml:"public_key"`
	SellerID   string `toml:"seller_id"`
	NotifyPath string `toml:"notify_path"`
}

type Config struct {
	Mode string `toml:"mode"`

	HTTPAddress string `toml:"http_listen_address"`
	GRPCAddress string `toml:"grpc_listen_address"`
	NotifyHost  string `toml:"http_notify_host"`

	AliPayment AliPay `toml:"alipay"`
	WXPayment  WxPay  `toml:"wxpay"`
}

func init() {
	flag.StringVar(&configFile, "conf", "./config.toml", "config file path")
	flag.Parse()
	log.SetFormatter(&log.JSONFormatter{})
	mc := client.NewLocalSmartConf(filepath.Dir(configFile))
	if err := mc.LoadObject(filepath.Base(configFile), &GlobalConfig); err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
				"file":  configFile,
			}).Fatal("load configure failed")
	} else {
		log.WithField("GlobalConfig", GlobalConfig).Info("load configure OK")
	}
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
