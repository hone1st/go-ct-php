package template

// {$namespace} 命名空间
// {$name}   名字
// {$ucName}   首字母小写

const Controller = `<?php

namespace {$namespace};

use App\Domain\Common\Controllers\BaseController;
use Illuminate\Http\Request;
use {$namespace_parent}\Services\{$name}Services;

class {$name}Controller extends BaseController
{
	/**@var {$name}Services ${$ucName}Services*/
    protected ${$ucName}Services;

    /**
     * @param {$name}Services ${$ucName}Services
     */
    public function __construct({$name}Services ${$ucName}Services)
    {
        $this->{$ucName}Services = ${$ucName}Services;
    }
	
    /**
     * 获取列表数据
     * @param Request $request
     */
    public function getList(Request $request)
    {
 		$input  = $request->input();
        $result = $this->{$ucName}Services->getList($input);
		return $this->success($result);
    }

    /**
     * 详情
     * @param Request $request
     */
    public function detail(Request $request)
    {
		$input  = $request->input();
        $result = $this->{$ucName}Services->detail($input);
        return $this->success($result);
    }

    /**
     * 添加
     * @param Request $request
     */
    public function add(Request $request)
    {
        $input  = $request->input();
        $result = $this->{$ucName}Services->add($input);
        return $this->success($result);
    }

    /**
     * 编辑
     * @param Request $request
     */
    public function edit(Request $request)
    {
        $input  = $request->input();
        $result = $this->{$ucName}Services->edit($input);
        return $this->success($result);
    }

    /**
     * 删除
     * @param Request $request
     */
    public function delete(Request $request)
    {
        $input  = $request->input();
        $result = $this->{$ucName}Services->delete($input);
        return $this->success($result);
    }
}
`
