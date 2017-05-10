package wxpay

import (
	"bytes"
	"encoding/xml"
	"git.chongdonglvxing.com/bz_platform/payment/util"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	AppId      string
	APIAddress string
	SellerID   string
	NotifyURL  string
	Key        string
}

type UnifyOrderReq struct {
	AppId          string `xml:"appid"`
	Body           string `xml:"body"`
	MchId          string `xml:"mch_id"`
	NonceStr       string `xml:"nonce_str"`
	NotifyUrl      string `xml:"notify_url"`
	TradeType      string `xml:"trade_type"`
	OpenID         string `xml:"openid"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
	TotalFee       int    `xml:"total_fee"`
	OutTradeNo     string `xml:"out_trade_no"`
	Sign           string `xml:"sign"`
}

type UnifyOrderResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppId      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	PrepayId   string `xml:"prepay_id"`
	TradeType  string `xml:"trade_type"`
}

func New(appId, apiAddress, sellerID, notifyURL, key string) *Client {
	return &Client{
		AppId:      appId,
		APIAddress: apiAddress,
		SellerID:   sellerID,
		NotifyURL:  notifyURL,
		Key:        key,
	}
}

func (c *Client) Pay(body string, tradeType string, openId string, spbillCreateIp string, totalFee int, outTradeNO string) (resp *UnifyOrderResp, err error) {
	var yourReq UnifyOrderReq
	yourReq.AppId = c.AppId
	yourReq.Body = body
	yourReq.MchId = c.SellerID
	yourReq.NonceStr = util.RandStringRunes(30)
	yourReq.NotifyUrl = c.NotifyURL
	yourReq.TradeType = tradeType
	yourReq.OpenID = openId
	yourReq.SpbillCreateIp = spbillCreateIp
	yourReq.TotalFee = totalFee
	yourReq.OutTradeNo = outTradeNO

	m := make(map[string]interface{}, 0)
	m["appid"] = yourReq.AppId
	m["body"] = yourReq.Body
	m["mch_id"] = yourReq.MchId
	m["notify_url"] = yourReq.NotifyUrl
	m["trade_type"] = yourReq.TradeType
	m["spbill_create_ip"] = yourReq.SpbillCreateIp
	m["total_fee"] = yourReq.TotalFee
	m["out_trade_no"] = yourReq.OutTradeNo
	m["nonce_str"] = yourReq.NonceStr
	yourReq.Sign = CalcSign(m, c.Key) //这个是计算wxpay签名的函数上面已贴出

	bytes_req, err := xml.Marshal(yourReq)
	if err != nil {
		return nil, err
	}

	str_req := string(bytes_req)
	//wxpay的unifiedorder接口需要http body中xmldoc的根节点是<xml></xml>这种，所以这里需要replace一下
	str_req = strings.Replace(str_req, "UnifyOrderReq", "xml", -1)
	bytes_req = []byte(str_req)

	//发送unified order请求.
	req, err := http.NewRequest("POST", c.APIAddress, bytes.NewReader(bytes_req))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/xml")
	//这里的http header的设置是必须设置的.
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")

	httpClient := http.Client{}
	httpResp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	httpBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	var unifyOrderResp UnifyOrderResp
	if err = xml.Unmarshal(httpBody, &unifyOrderResp); err != nil {
		return nil, err
	}

	return &unifyOrderResp, nil
}
