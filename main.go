package main

import (
	"fmt"
	"git.chongdonglvxing.com/bz_platform/payment/alipay"
	"git.chongdonglvxing.com/bz_platform/payment/config"
	"git.chongdonglvxing.com/bz_platform/payment/handler"
	api "git.chongdonglvxing.com/bz_platform/payment/pbgen"
	"git.chongdonglvxing.com/bz_platform/payment/wxpay"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	aliClient := alipay.New(config.GlobalConfig.AliPayment.APIID,
		fmt.Sprintf("%s%s", config.GlobalConfig.NotifyHost, config.GlobalConfig.AliPayment.NotifyPath),
		[]byte(config.GlobalConfig.AliPayment.PublicKey),
		[]byte(config.GlobalConfig.AliPayment.PrivateKey),
		true)

	wxClient := wxpay.New(config.GlobalConfig.WXPayment.APIID,
		config.GlobalConfig.WXPayment.APIAddress,
		config.GlobalConfig.WXPayment.SellerID,
		fmt.Sprintf("%s%s", config.GlobalConfig.NotifyHost, config.GlobalConfig.WXPayment.NotifyPath),
		config.GlobalConfig.WXPayment.Key)

	go launchHTTPServer(config.GlobalConfig.HTTPAddress, aliClient, wxClient)
	go launchGrpcServer(config.GlobalConfig.GRPCAddress, handler.NewPaymentHandler(wxClient, aliClient))
	wg.Wait()
}

func launchHTTPServer(addr string, aliClient *alipay.AliPay, wxClient *wxpay.Client) {
	http.HandleFunc(config.GlobalConfig.WXPayment.NotifyPath, wxClient.Callback)
	http.HandleFunc(config.GlobalConfig.AliPayment.NotifyPath, aliClient.Callback)
	http.ListenAndServe(addr, nil)
	wg.Done()
}

func launchGrpcServer(addr string, h api.PaymentServer) {
	grpcServer := grpc.NewServer()
	api.RegisterPaymentServer(grpcServer, h)

	lsner, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	grpcServer.Serve(lsner)
	wg.Done()
}
