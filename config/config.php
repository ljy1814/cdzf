<?php

/**
 * Created by PhpStorm.
 * User: yan
 * Date: 2017/5/16
 * Time: 下午9:19
 */
class Config
{
    private static $configs = [
        "alipay" => [
            'partner' => '2088421965342291',
            'sign_type' => 'RSA',
            'input_charset' => 'utf-8',
            'app_id' => '2016101302148385',
            'seller_email' => 'cdlx2016@126.com',
            'key' => 'enujty3catmp7pb7zbwukwlbdsm6hj1h',
            'payment_type' => 1,
            'notify_url' => '',
            'return_url' => '',
            'transport' => 'http',
            'ali_public_key_path' => __DIR__ . '/alipay/ali_public_key.pem',
            'private_key_path' => __DIR__ . '/alipay/rsa_private_key.pem',
        ],
        "wechat" => [
            "baseUrl" => "http://pay.chongdonglvxing.com/pay/wechat/",
            'appid'=>"wx03c90d06969dc5f4",
            'wxmchid' => '1398841802',

            'appidapp'=>"wx65f31f9bdd5b376f",
            'wxmchidapp' => '1407166902',

            'wxpayapp_key' => 'T7VkBo6ZotBasPDGqhtGsBtG6XhatRfe',
            'wxpayapp_sslcert_path' => __DIR__.'/wechat/apiclient_cert.pem',
            'wxpayapp_sslkey_path' => __DIR__.'/wechat/apiclient_key.pem',

            'curl_timeout' => 30,
            'REPORT_LEVENL' => 0,//错误级别是否上报
            'notify_url' => 'http://cd.loserpm.com/api/wechat/notify',
            'timeout'=>6000
        ],
        "server_addr" => "0.0.0.0",
        "server_port" => 9002,
        "debug"=>true

    ];


    public static function Get(string $key, $default = null)
    {
        $keys = explode('.', $key);
        if (isset($keys[1])) {
            return self::$configs[$keys[0]][$keys[1]] ?? self::$configs[$key] ?? $default ?? "";
        }
        return self::$configs[$key] ?? $default ?? "";
    }

}