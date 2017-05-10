package handler

import (
	"git.chongdonglvxing.com/bz_platform/payment/alipay"
	api "git.chongdonglvxing.com/bz_platform/payment/pbgen"
	"git.chongdonglvxing.com/bz_platform/payment/wxpay"
	"golang.org/x/net/context"
)

type PaymentHandler struct {
	aliClient *alipay.AliPay
	wxClient  *wxpay.Client
}

func NewPaymentHandler(wc *wxpay.Client, ac *alipay.AliPay) *PaymentHandler {
	ph := &PaymentHandler{
		aliClient: ac,
		wxClient:  wc,
	}

	return ph
}

func (h *PaymentHandler) AliPay(ctx context.Context, in *api.AliPayRequest) (out *api.AliPayResponse, err error) {
	var p = alipay.AliPayTradeWapPay{}
	//	p.NotifyURL = h.aliClient
	p.ReturnURL = in.ReturnURL
	p.Subject = in.Subject
	p.OutTradeNo = in.OutTradeNo
	p.TotalAmount = in.TotalAmount
	p.ProductCode = in.ProductCode

	out = &api.AliPayResponse{}
	url, err := h.aliClient.TradeWapPay(p)
	if err != nil {
		out.Error = &api.Error{
			Msg: err.Error(),
		}
		return out, nil
	}

	out.Url = url.String()

	return out, nil
}

func (h *PaymentHandler) WxPay(ctx context.Context, in *api.WxPayRequest) (out *api.WxPayResponse, err error) {
	out = &api.WxPayResponse{}
	resp, err := h.wxClient.Pay(in.Body, in.TradeType, in.OpenId, in.SpbillCreateIp, int(in.TotalFee), in.OutTradeNO)
	if err != nil {
		out.Error = &api.Error{
			Msg: err.Error(),
		}
		return out, nil
	}
	/*
		string ReturnCode  = 1;
		string ReturnMsg   = 2;

		string AppId        = 3;
		string MchId       = 4;
		string NonceStr    = 5;
		string Sign         = 6;
		string ResultCode  = 7;
		string PrepayId    = 8;
		string TradeType   = 9;
	*/
	out.ResultCode = resp.ResultCode
	out.ReturnMsg = resp.ReturnMsg
	out.AppId = resp.AppId

	out.AppId = resp.AppId
	out.MchId = resp.MchId
	out.NonceStr = resp.NonceStr
	out.Sign = resp.Sign
	out.ResultCode = resp.ResultCode
	out.PrepayId = resp.PrepayId
	out.TradeType = resp.TradeType

	if out.ReturnCode == "FAIL" {
		out.Error = &api.Error{
			Msg: out.ReturnMsg,
		}
		return out, nil
	}

	if out.ResultCode == "FAIL" {
		out.Error = &api.Error{
			Msg: "Result Code Failed",
		}
		return out, nil
	}

	return out, nil
}
