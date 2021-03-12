package admin_template

// {$module}
// {$model}

const Service = `<?php


namespace app\{$module}\service;


use app\{$module}\model\{$model} as {$model}Model;
use app\common\service\BaseService;

class {$model} extends BaseService {

	public function index(int $id = 0, int $page = 1, int $limit = 10) {
		$whereMap = [];
        if($id != 0) {
			array_push($whereMap, ['id', '=', $id]);
        }
        $count = (int){$model}Model::where($whereMap)->count();
        $data  = {$model}Model::order('id desc')
				->where($whereMap)
				->page($page, $limit)
                ->select();
        return $this->makePage($data, $count, $page, $limit);
    }

    public function update(int $id, array $data) {
        return {$model}Model::where(['id' => $id])->update($data);
    }

    public function create(array $data) {
        return {$model}Model::create($data);
    }

    public function delete(int $id) {
        return {$model}Model::where(['id' => $id])->delete();
    }
}`
