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

    /**
     * 详情
     * @param Request $request
     */
    public function detail(Request $request)
    {
        $result = $this->{$ucName}Repository->detail($request->input());
        return $this->success($result);
    }

    /**
     * 添加
     * @param Request $request
     */
    public function add(Request $request)
    {
        $result = $this->{$ucName}Repository->add($request->input());
        return $this->success($result);
    }

    /**
     * 编辑
     * @param Request $request
     */
    public function edit(Request $request)
    {
        $result = $this->{$ucName}Repository->edit($request->input());
        return $this->success($result);
    }

    /**
     * 删除
     * @param Request $request
     */
    public function delete(Request $request)
    {
        $result = $this->{$ucName}Repository->delete($request->input());
        return $this->success($result);
    }
}
`
