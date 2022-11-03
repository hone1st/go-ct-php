package template

// {$namespace} 命名空间
// {$name}   	名字
// {$name_id}   表名字+_id 名字

const Repository = `<?php

namespace {$namespace};


class {$name}Repository 
{
	
	/**
     * 列表
     * @param array $params
     */
    public function getList(array $params)
    {
        return {$name}::query()
{$when}
			->paginate($params['page_size'] ?? 15);
    }

	/**
     * 详情
     * @param array $params
     */
    public function detail(array $params)
    {
        return {$name}::query()
{$when}
				->first();
    }


	/**
     * 编辑
     * @param array $input
     */
    public function edit(array $input)
    {
        return {$name}::query()->where('id', $input['{$name_id}'])->update([
			{$fields_map}
        ]);
    }

	/**
     * 删除
     * @param array $input
     */
    public function delete(array $input)
    {
        return {$name}::query()->where('id', $input['{$name_id}'])->delete();
    }

	/**
	 * 新增
     * @param array $input
     */
    public function add(array $input)
    {
        return {$name}::query()->create([
            {$fields_map}
        ]);
    }

}`
