<?php

/**
 * Created by PhpStorm.
 * User: yan
 * Date: 2017/5/23
 * Time: 上午8:26
 */
class Logs
{
    public function Set()
    {
        array_map(function ($x) {
            if(!is_string($x)){
                $x = json_encode($x,JSON_UNESCAPED_UNICODE);
            }
            $url = $_SERVER['REQUEST_URI'] ?? '';
            error_log('['.date('Y-m-d H:i:s').']'.$url." ".$x .' '."\n", 3, __DIR__."/../logs/debug.log");
        }, func_get_args());

        die(1);
    }
}
