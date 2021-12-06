package template

// {$module}
// {$model}

const AdminService = `<?php


namespace app\{$module}\service;


use app\{$module}\model\{$model} as {$model}Model;
use app\common\service\BaseService;

class {$model} extends BaseService {

    /**
     * 列表数据
     * @param int      $id   主键的id
     * @param array    $page   当前页数
     * @param array    $limit   每页的数量
     * @return bool
     */
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

    /**
     * 更新数据
     * @param int      $id   主键的id
     * @param array    $data   数据数组
     * @return bool
     */
    public function update(int $id, array $data) {
        return (new {$model}Model)->allowField(true)->save($data, ['id' => $id]);
    }


    /**
     * 创建数据
     * @param array    $data   数据数组
     * @return bool
     */
    public function create(array $data) {
        return {$model}Model::create($data);
    }

   /**
     * 删除
     * @param int $id      主键的id
     */
    public function delete(int $id) {
        return {$model}Model::where(['id' => $id])->delete();
    }

   /**
     * 审核通过
     * @param int $id      主键的id
     * @param int $auditId 审核人
     */
    public function pass(int $id, int $auditId)
    {
       // 做审核通过的操作
    }
	
    /**
     * 审核不通过
     * @param int    $adminId   审核人
     * @param string $rejectReason 拒绝原因
     * @param int    $id  主键的id
     * @return bool
     */
	public function noPass(int $adminId, string $rejectReason, int $id)
    {
		// 做审核拒绝的操作
        return false;
    }
}`
