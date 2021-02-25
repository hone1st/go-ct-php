package admin_template

// {$property} 表示数据库的字段参数
// {$model} 表示大写的模型名称
// {$module}
const Model = `<?php

namespace app\{$module}\model;

use app\group\append\AppendAny;
use think\Model;

{$property}
class {$model} extends Model {
    use AppendAny;
	protected $append_any = 'text';
}
`
