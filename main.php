<?php
/**
 * Created by PhpStorm.
 * User: yan
 * Date: 2017/5/16
 * Time: 下午8:50
 */
require __DIR__ . "/bootstrap.php";
$http = new swoole_http_server(Config::Get('server_addr'), Config::Get('server_port'));

$http->on('request', function (swoole_http_request $request, swoole_http_response $response) {
    $app = new \App\App($request, $response);
    $app->router();

});

$http->set([
    'worker_num' => 10,
    'log_level' => 0,
    'pid_file' => __DIR__.'/logs/server.pid',
    'log_file'=>__DIR__.'/logs/app.log',
    'daemonize'=> false
]);
if(!file_exists(__DIR__."/logs")){
    mkdir(__DIR__.'/logs');
}
$http->start();