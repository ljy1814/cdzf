package client

import (
	"fmt"
	api "git.chongdonglvxing.com/bz_platform/payment/pbgen"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
)

var (
	address = "localhost:9090"
)

func TestCreateDeleteUser(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	clnt := api.NewPaymentClient(conn)

	resp, err := clnt.WxPay(context.Background(), &api.WxPayRequest{
		Body:           "微信支付body",
		TradeType:      "NATIVE",
		SpbillCreateIp: "127.0.0.1",
		TotalFee:       1,
		OutTradeNO:     "12314141",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", resp)
}
