<?php
/**
 * Created by PhpStorm.
 * User: yan
 * Date: 2017/5/17
 * Time: 上午9:26
 */
echo '[1;34;0m______________________________________________________________________[0m'."\n";
echo '[1;34;0m|        时间        |方法|状态|        PATH        |        IP        |[0m'."\n";
echo '[1;34;0m----------------------------------------------------------------------[0m'."\n";
require __DIR__."/config/config.php";
require __DIR__."/alipay/AopSdk.php";
require __DIR__."/wechat/WxPay.JsApiPay.php";
require __DIR__."/app/app.php";

function dd()
{
    array_map(function ($x) {
        var_dump($x);
    }, func_get_args());

}
function debug()
{
    array_map(function ($x) {
        if(!is_string($x)){
            $x = json_encode($x,JSON_UNESCAPED_UNICODE);
        }
        $url = $_SERVER['REQUEST_URI'] ?? '';
        error_log('['.date('Y-m-d H:i:s').']'.$url." ".$x .' '."\n", 3, __DIR__."/logs/debug.log");
    }, func_get_args());
}