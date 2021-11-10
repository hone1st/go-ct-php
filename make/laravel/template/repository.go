package template

// {$namespace} 命名空间
// {$name}   名字

const Repository = `<?php

namespace {$namespace};


class {$name}Repository 
{
	
    /**
     * @param array $params
     * @return array|object[]
     */
    public function getList(array $params)
    {
        // do somethings
        return [];
    }

}`
