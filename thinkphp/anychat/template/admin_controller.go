package template

// {$lModel} 首字母小写的模型名字
// {$model} 首字母大小的模型名字
// {$module} 模块名字

const AdminController = `<?php


namespace app\{$module}\admin;


use app\{$module}\service\{$model} as {$model}Service;
use app\{$module}\model\{$model} as {$model}Model;
use app\common\base\AdminBaseV1;
use app\common\facade\ShowView;
use think\App;

// 模板代码
class {$model} extends AdminBaseV1 {
	
	protected $hisiModel = '{$model}';

    protected $autoValidate = [
            'index.get'     => \app\{$module}\validate\{$model}::class . '.index',
            'del.get' => \app\{$module}\validate\{$model}::class . '.delete',
            'update.post'    => \app\{$module}\validate\{$model}::class . '.update',
            'create.post'   => \app\{$module}\validate\{$model}::class . '.create',
    ];

    /**
     * @var {$model}Service
     */
    private ${$lModel}Service;

    /**@inheritdoc */
    public function __construct({$model}Service ${$lModel}Service, App $app = null) {
        parent::__construct($app);
        $this->{$lModel}Service = ${$lModel}Service;
    }


    /**
     * 列表数据
     */
    public function index() {
        // 渲染数据
        if (!$this->request->isAjax()) {
			// 下拉选择例子
			$this->assign('select_cn', [0 => '男', 1=> '女', 2=> '性别不详']);
            return $this->fetch();
        } else {
            $page  = $this->request->param('page');
            $limit = $this->request->param('limit');
            $id    = $this->request->param('id/d', 0);
			// 下拉选择例子
            $selectExample    = $this->request->param('select_example/d', 0);
            return ShowView::data($this->{$lModel}Service->index($id, $page, $limit))->send();
        }
    }


    /**
     * 更新数据
     */
    public function create() {
        if($this->request->isPost()) {
            if($this->{$lModel}Service->create($this->request->post())) {
				 $this->success('添加成功');
            }
             $this->error('添加失败');
        }
		$this->assign('radio_example', [0 => '男', 1=> '女', 2=> '性别不详']);
		$this->assign('content_example', '<p>123123123131</p>');
		$this->assign('select_example', [0 => '男', 1=> '女', 2=> '性别不详']);
		return $this->fetch();
    }

    // 更新
    public function update() {
        $id = $this->request->param('id');
        // 渲染数据
        if (!$this->request->isAjax() && $this->request->isGet() ) {
 			$item = {$model}Model::where(['id' => $id])->find();
            $this->assign('id', $id);
			$item['radio_example'] = 0;
			$item['select_example'] = 0;
			$item['content_example'] = '<p>123123123131</p>';
            $this->assign('item', $item);
			$this->assign('radio_example', [0 => '男', 1=> '女', 2=> '性别不详']);
			$this->assign('select_example', [0 => '男', 1=> '女', 2=> '性别不详']);
            return $this->fetch();
        } elseif($this->request->post()) {
            if ($this->{$lModel}Service->update($id, $this->request->post())) {
                  $this->success('修改成功');
            }
             $this->error('修改失败');
        }
    }

	 /**
     *  审核通过
     */
    public function pass()
    {
        if ($this->request->isAjax()) {
            $id    = $this->request->request('id');
            $login = session('admin_user');
            $this->{$lModel}Service->pass($id, $login['uid']);
            $this->success('审核成功!');
        }
    }
	
	/**
	 * 审核不通过	
	*/

    public function noPass()
    {
        $id = $this->request->param('id/d', 0);
        if ($this->request->isPost()) {
            $data = $this->request->post();
            if ($id == 0) {
                $this->error('修改失败');
            }
            $login = session('admin_user');
            if (!$this->{$lModel}Service->noPass($login['uid'], $data['reject_reason'], $id)) {
                $this->error('修改失败');
            }
            $this->success('修改成功');

        } else {
            $info = {$model}Model::where(['id' => $id])->allowEmpty()->findOrEmpty();
            $this->assign('info', $info);

            return $this->fetch();
        }
    }
}`