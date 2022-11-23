package template

// {$namespace} 命名空间
// {$namespace_parent} 上一级的命名空间
// {$name}   名字
// {$ucName}   首字母小写

const Services = `<?php

namespace {$namespace};


use {$namespace_parent}\Repository\{$name}Repository;

class {$name}Services
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
     * 获取限制数量
     * @param array $params
     */
    public function getList(array $params)
    {
        return $this->{$ucName}Repository->getList($params);
    }

    /**
     * 列表自动分页
     * @param array $params
     */
    public function getPaginate(array $params)
    {
        return $this->{$ucName}Repository->getPaginate($params);
    }

    /**
     * 详情
     * @param array $params
     */
    public function detail(array $params)
    {
        return $this->{$ucName}Repository->detail($params);
    }

    /**
     * 添加
     * @param array $input
     */
    public function add(array $input)
    {
        return $this->{$ucName}Repository->add($input);
    }

    /**
     * 编辑
     * @param array $input
     */
    public function edit(array $input)
    {
        return $this->{$ucName}Repository->edit($input);
    }

    /**
     * 删除
     * @param array $input
     */
    public function delete(array $input)
    {
        return $this->{$ucName}Repository->delete($input);
    }
}
`
