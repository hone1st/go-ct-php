package template

// {$namespace} 命名空间
// {$name}   名字
// {$ucName}   首字母小写

const Controller = `<?php

namespace {$namespace};


use Illuminate\Http\Request;
use App\Repository\{$name}Repository;

class {$name}Controller extends BaseController
{
	/**@var {$name}Repository ${$ucName}Repository*/
    protected ${$ucName}Repository;

    /**
     * @param {$name}Repository ${$ucName}Repository
     */
    public function __construct({$name}Repository ${$ucName}Repository)
    {
        $this->{$ucName}Repository = ${$ucName}Repository;
    }
	
    /**
     * 获取列表数据
     * @param Request $request
     */
    public function getList(Request $request)
    {
        $result = $this->{$ucName}Repository->getList($request->input());
		return $this->success($result);
    }
}
`
