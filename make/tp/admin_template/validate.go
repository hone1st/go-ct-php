package admin_template

// {$model}
// {$module}
// {$message} 表示字段映射注释 'id'=> '主键'
// {$rule} 表示字段映射中括号 'id' => [],'xxx' => [],
// {$scene} 表示字段 'id','xx','xxx'

const Validate = `<?php


namespace app\{$module}\validate;


use think\Validate;
use app\common\validate\PageValidate;

class {$model} extends PageValidate {

    /**@inheritdoc */
    protected $rule = [
            'limit' => ['number', 'integer'],
            'page'  => ['number', 'integer'],
{$rule}
    ];

    /**@inheritdoc */
    protected $message = [
            'limit' => '每页显示的数量',
            'page'  => '当前页',
{$message}
    ];

    /**@inheritdoc */
    protected $scene = [
            'index'  => ['limit', 'page'],
            'create' => [{$scene}],
            'update' => [{$scene}],
            'delete' => ['id'],
    ];
}`
