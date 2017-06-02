<?php
/**
 * Created by PhpStorm.
 * User: yan
 * Date: 2017/5/16
 * Time: 下午8:50
 */
require __DIR__ . "/bootstrap.php";
function serverStart($daemonize = false,$worker_num = 0)
{
    if(!$daemonize){
        echo '[1;34;0m______________________________________________________________________[0m'."\n";
        echo '[1;34;0m|        时间        |方法|状态|        PATH        |        IP        |[0m'."\n";
        echo '[1;34;0m----------------------------------------------------------------------[0m'."\n";
    }else{
        print "server starting\n";
    }

    $http = new swoole_http_server(Config::Get('server_addr'), Config::Get('server_port'));

    $http->on('request', function (swoole_http_request $request, swoole_http_response $response) {
        $app = new \App\App($request, $response);
        $app->router();

    });

    $http->set([
        'worker_num' => $worker_num,
        'log_level' => 0,
        'pid_file' => __DIR__.'/logs/server.pid',
        'log_file'=>__DIR__.'/logs/app.log',
        'daemonize'=> $daemonize
    ]);
    if(!file_exists(__DIR__."/logs")){
        mkdir(__DIR__.'/logs');
    }
    $http->start();
}
function serverStop()
{
    print "server finish ";
    $pid = file_get_contents(__DIR__."/logs/server.pid");
    $p = swoole_process::kill($pid, 0);
    if ($p){
        swoole_process::kill($pid, 15);
    }
    print ".";
    sleep(1);
    print ".";
    sleep(1);
    print ".\n";
}