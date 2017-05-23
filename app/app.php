<?php

namespace App;

/**
 * Created by PhpStorm.
 * User: yan
 * Date: 2017/5/17
 * Time: 上午10:20
 */
class App
{
    private $request;
    private $response;
    private $routing = [
        "GET/alipay" => "ALiPay",
        "GET/wechat" => "WeChat",
        "GET/MP_verify_vStFdbnWe61H04fE.txt" => "Token"
    ];
    private $status = 200;
    private $end = "";
    public function __construct(\swoole_http_request $request, \swoole_http_response $response)
    {

        $this->request = $request;
        $this->response = $response;

    }
    public function __destruct()
    {
        // TODO: Implement __destruct() method.
        $this->response->status($this->status);
        if(\Config::Get('debug')){
            $color = 31;
            if($this->status == 200){
                $color = 32;
            }
            echo '|'.date('Y-m-d H:i:s').'|'."\033[1;33;0m".$this->request->server['request_method']."\e[0m| \033[1;$color;0m".$this->status."\033[0m| ".sprintf('% 18s', $this->request->server['path_info']).'  |'.sprintf('% 15s', $this->request->server['remote_addr'])."  | \n";
            echo '[1;34;0m----------------------------------------------------------------------[0m'."\n";
        }
        $this->response->end($this->end);
    }

    public function router()
    {
        $this->response->header("Content-Type", "text/json; charset=utf-8");
        $this->response->header("Server", "ChongDongServer");
        $path = $this->request->server['request_method'] . $this->request->server['path_info'];
        if (!isset($this->routing[$path])) {

            $this->status = 404;
            $this->end = json_encode(['code'=>404,'err_msg'=>"很抱歉，您要访问的页面不存在！"]);
            return;
        }
        call_user_func(['App\App', $this->routing[$path]]);

    }

    public function ALiPay()
    {
        $alipay_config = \Config::Get('alipay');
        $paraToken = array(
            "out_trade_no" => mt_rand(),
            "seller_id" => $alipay_config['partner'],
            "total_amount" => 0.01,
            "subject" => mt_rand(),
            "body" => mt_rand(),
        );
        $aop = new \AopClient();
        $aop->gatewayUrl = 'https://openapi.alipay.com/gateway.do';
        $aop->appId = $alipay_config['app_id'];
        $aop->rsaPrivateKeyFilePath = $alipay_config['private_key_path'];
        $aop->alipayPublicKey = $alipay_config['ali_public_key_path'];
        $aop->apiVersion = '1.0';
        $aop->postCharset = 'UTF-8';
        $aop->format = 'json';

        $request = new \AlipayTradeWapPayRequest();
        $request->setNotifyUrl($alipay_config['notify_url']);
        $request->setReturnUrl($alipay_config['return_url']);
        $request->setBizContent(json_encode($paraToken));
        $result = $aop->pageExecute($request);

        $this->end = json_encode(['code'=>200,"data"=>$result]);
    }

    public function WeChat()
    {
        $tools = new \JsApiPay();
        $code = $this->request->get['code'] ?? "";
        //通过code获得openid
        if ($code == "") {
            //触发微信返回code码
            $baseUrl = urlencode(\Config::Get("wechat.baseUrl"));
            $url = $tools->CreateOauthUrlForCode($baseUrl);
            $this->status = 301;
            $this->response->header('location',$url);
           // $this->response->end(json_encode(['code'=>301,'data'=>$url]));
            return;
        }

        //获取code码，以获取openid

        $openid = $tools->getOpenidFromMp($code);
        $input = new \WxPayUnifiedOrder();
        $input->SetBody("test");
        $input->SetAttach("test");
        $input->SetOut_trade_no(\WxPayConfig::MCHID . date("YmdHis"));
        $input->SetTotal_fee("1");
        $input->SetTime_start(date("YmdHis"));
        $input->SetTime_expire(date("YmdHis", time() + 6000));
        $input->SetGoods_tag("test");
        $input->SetNotify_url("http://paysdk.weixin.qq.com/example/notify.php");
        $input->SetTrade_type("JSAPI");
        $input->SetOpenid($openid);
        $order = \WxPayApi::unifiedOrder($input);
        $jsApiParameters = $tools->GetJsApiParameters($order);
        $this->end = $jsApiParameters;
    }

    public function Token()
    {
        $this->response->end("vStFdbnWe61H04fE");
    }
}