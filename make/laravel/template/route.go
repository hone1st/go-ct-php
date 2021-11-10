package template

// {$prefix} 默认是控制器的前缀  RefundController => refund   OrderRefundController => order_refund
// {$namespace} 命名空间
// {$controller} 控制器名字

const Route = `<?php

use Illuminate\Support\Facades\Route;
use {$namespace}\{$controller}Controller;

Route::group([
    'prefix'     => '{$prefix}',
    'middleware' => [
        // 中间件
    ],
], function () {
	// 获取列表数据
    Route::get('getList', [{$controller}Controller::class, 'getList']);
});`
