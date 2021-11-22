package template

// {$namespace} 命名空间
// {$name}   名字

const Repository = `<?php

namespace {$namespace};


class {$name}Repository 
{
	
	/**
     * @param array $params
     */
    public function getList(array $params)
    {
        return {$name}::query()->paginate($params['page_size'] ?? 15);
    }

    public function edit(array $input)
    {
        return {$name}::query()->where('id', $input['{$name_id}'])->update([
			{$fields_map}
        ]);
    }

    public function delete(array $input)
    {
        return {$name}::query()->where('id', $input['{$name_id}'])->delete();
    }

    public function add(array $input)
    {
        return {$name}::query()->create([
            {$fields_map}
        ]);
    }

}`
