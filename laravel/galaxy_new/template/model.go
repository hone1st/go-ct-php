package template

// {$property} 数据库表的字段类型和注释
// {$namespace} 命名空间
// {$table} 表名字
//  {$fields} 数据库字段的列表
//  {$fieldsTrans} 数据库字段对应的类型转换器

//  {$fieldsRule} 数据库字段对应验证 'deleted_at' => 'nullable', 'created_at' => 'nullable', 'updated_at
//  默认其他字段是必须的 如果是字符串的换就限制长度string int类型就限定integer  required

const Model = `<?php

namespace {$namespace};

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;


{$property}
class {$model} extends Model
{
    use SoftDeletes;

    public $table = '{$table}';

	protected $connection = '{$connect}';

    const CREATED_AT = 'created_at';
    const UPDATED_AT = 'updated_at';

    protected $dates = ['deleted_at'];

	protected $casts = [
{$fieldsTrans}
	];

    public $fillable = [
{$fields}
    ];
}
`
