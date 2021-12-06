package template

// {$cols} =>  {field: 'id', title: '主键', width: 100},

const AdminIndex = `<!-- 表格上方的搜索区域 -->
<div class="table_top">
 例子(搜索id)：
	<form class="layui-form layui-form-pane" lay-filter="search_data">
		<div class="layui-input-inline">
            <input class="layui-input" name="id"  placeholder="id">
        </div>
		<div class="layui-input-inline">
           <select name="select_example">
				<option value="-1">全部</option>
				{foreach $select_cn as $key=>$vo }
					<option value="{$key}">{$vo}</option>
				{/foreach}
			</select>
        </div>
		<div class="layui-input-inline">
			<div class="layui-input-inline">
				<a class="layui-btn layui-btn-normal" id="search" data-type="search">搜索</a>
				<a class="layui-btn layui-btn-normal" id="reset" data-type="reset"> 重置</a>
			</div>
		</div>
	</form>
</div>
<!--表格数据的渲染-->
<table id="table_data" lay-filter="table_data"></table>
{include file="system@block/layui" /}
<!--表格的行尾的按钮-->
<script type="text/html" title="" id="table_row_btn">
	 <a href="{:url('update')}?id={{d.id}}" class="layui-btn layui-btn-xs layui-btn-normal hisi-iframe"
       hisi-data="{width: '800px', height: '800px'}"
       title="编辑数据">编辑数据</a>
	<a href="{:url('del')}?id={{ d.id }}" class="layui-btn layui-btn-xs layui-btn-danger j-tr-del">删除数据</a>

	<!--审核操作 无延迟刷新-->
    <a href="{:url('pass')}?id={{ d.id }}" confirm="确定通过吗？" class="layui-btn layui-btn-xs layui-btn-normal hisi-ajax"  custom_refresh="true">通过</a>
    <a href="{:url('noPass')}?id={{d.id}}" class="layui-btn layui-btn-xs layui-btn-normal hisi-iframe"
       hisi-data="{width: '400px', height: '400px'}"
       title="拒绝通过ID：【{{d.id}}】">拒绝</a>

</script>
<!--表格的嵌入的头部的工具栏-->
<script type="text/html" id="table_top_tool">
    <div class="layui-btn-group fl">
        <a href="{:url('create')}"
           class="layui-btn layui-btn-primary layui-btn-sm layui-icon layui-icon-add-circle-fine hisi-iframe"
           hisi-data="{width: '800px', height: '800px'}" title="添加数据">&nbsp;添加数据</a>
    </div>
</script>
<script type="text/javascript">
    layui.use(['table','form'], function () {
        // 显示table
        let table = layui.table, $ = layui.jquery,form = layui.form;

		// 初始化搜索条件
		let reset = {id: 0, select_example: -1}
        table.render({
            elem: '#table_data'
            , url: '{:url()}' //数据接口
            , parseData: function (res) { //res 即为原始返回的数据
                return {
                    "code": res.recode, //解析接口状态
                    "msg": res.msg, //解析提示文本
                    "count": res.data.count, //解析数据长度
                    "data": res.data.data //解析数据列表
                };
            }
			,cellMinWidth: 100
            // , page: true //开启分页
            , skin: 'row'
            , id: 'id'
            , even: true
            , limit: 20
            , text: {none: '暂无相关数据'}
            , toolbar: '#table_top_tool'
            , defaultToolbar: ['filter']
            , cols: [[ //表头
				{$cols}
                {title: '管理操作', templet: '#table_row_btn'}
            ]]
        });

        let active = {
            search: function () {
                let search = form.val('search_data')
                //执行重载
                table.reload('id', {
                    url: '{:url()}',
                    page: {
                        curr: 1 //重新从第 1 页开始
                    }
                    , where: search
                });
            },
            reset:function () {
                form.val('search_data',reset)
                form.render()
            }
        };
		// 搜索
        $('#search').on('click', function () {
            let type = $(this).data('type');
            active[type] ? active[type].call(this) : '';
        });
		// 重置
		$('#reset').on('click', function () {
            let type = $(this).data('type');
            active[type] ? active[type].call(this) : '';
        });
    });

</script>

<!--使用例子     {field: 'images', title: '图片', align: 'center', templet: '#imagesTpl'},
                {field: 'video', title: '视频', align: 'center', templet: '#videoTpl'},
                -->
<!--图片模板-->
<script type="text/html" id="imagesTpl">
    {{#  if(d.cover != null  && d.cover != ''){ }}
    <div class="ssy_img_parent">
        <img src="{{ d.cover }}" width="40" class="ssy_img" height="40">
    </div>
    {{#  } }}
</script>
<!--视频模板-->
<script type="text/html" id="videoTpl">
	<!--    数据结构 video: {cover:url,path:url}-->
    {{#  if(d.video != null){ }}
    <li>
        <img src="{{ d.video.cover }}" width="50" height="50" onclick='playerVideo("{{ d.video.path }}")'>
    </li>
    {{#  } }}
</script>
<script>
    /**
     *  播放视频
     */
    function playerVideo(vUrl) {
        var loadstr = '<video width="100%" height="100%"  controls="controls" autobuffer="autobuffer"  autoplay="autoplay" loop="loop">' +
            '<source src=' + vUrl + ' type="video/mp4"></source></video>';
        layer.open({
            type: 1,
            title: false,
            area: ['730px', '500px'],
            shade: [0.8, 'rgb(14, 16, 22)'],
            skin: 'demo-class',
            content: loadstr
        });
    }


</script>
`

// {$update}
/**
	<div class="layui-form-item">
            <label class="layui-form-label">地区ID :</label>
            <div class="layui-input-block">
                <input type="text" class="layui-input field-name" name="id" lay-verify="required"
                       autocomplete="off" placeholder="地区ID">
            </div>
     </div>
*/

const NoPass = `<form class="layui-form" action="{:url()}" method="post" id="editForm">

  <div class="layui-form-item">
    <label class="layui-form-label">审核备注</label>
    <div class="layui-input-inline">
      <textarea class="layui-textarea field-audit_remark" name="reject_reason" style="width: 100%;" rows="3" cols="20"></textarea>
    </div>
  </div>
  <div class="layui-form-item">
    <div class="layui-input-block">
      <input type="hidden" class="field-id" name="id" value="{$info['id']}">
    </div>
  </div>
  <div class="pop-bottom-bar">
    <button type="submit" class="layui-btn layui-btn-normal" lay-submit="" lay-filter="formSubmit" hisi-data="{pop: true,custom_refresh: true}">提交保存</button>
    <a href="{:url('index')}" class="layui-btn layui-btn-primary ml10"><i class="aicon ai-fanhui"></i>返回</a>
  </div>
</form>
{include file="system@block/layui" /}
<script>
  layui.use(['form', 'func'], function() {
    var $ = layui.jquery, form = layui.form;
    layui.func.assign({:json_encode($formData)});
    console.log({:json_encode($formData)})
  });
</script>`

const Update = `<form class="layui-form alignment layui-form-pane" action="{:url('')}?id={$id}" method="post">
<!--    隐藏的主键的id -->
	<input type="hidden" name="id" value="{$id}">
    <div class="form_update">
		{$update}

	    <!-- 上传图片的示例 -->
		<div class="layui-form-item">
            <label class="layui-form-label">上传图片示例</label>
            <div class="layui-input-block">
                <button type="button" class="layui-btn" id="img_btn">请选择头像</button>
				 <div class="ssy_img_parent">
						<img  width="40" class="ssy_img" height="40" id="img_show_example"/>
				</div>
            </div>
			<!-- 上传的图片的字段 -->
            <input type="hidden"  name="img_example" />
        </div>

		<!-- 单选的示例 -->
		<div class="layui-form-item">
            <label class="layui-form-label">单选的示例</label>
            <div class="layui-input-block">
                {foreach $radio_example as $key=>$vo }
				{if $item['radio_example'] == $key}
				<input type="radio" name="radio_example" value="{$key}" title="{$vo}" checked="">
				{else /}
				<input type="radio" name="radio_example" value="{$key}" title="{$vo}">
				{/if}
                {/foreach}
            </div>
        </div>
		
		<!-- 下拉选择的示例 -->
		<div class="layui-form-item">
            <label class="layui-form-label">下拉选择的示例</label>
            <div class="layui-input-block">
				<select name="select_example">
 				    {foreach $select_example as $key=>$vo }
					{if $key == $item['select_example']}
					<option value="{$key}" selected="">{$vo}</option>
					{else /}
					<option value="{$key}">{$vo}</option>
					{/if}
					{/foreach}
				</select>
            </div>
        </div>

		<!-- 富文本的例子 -->		
		 <div class="layui-form-item">
            <label class="layui-form-label">新闻内容</label>
            <div class="layui-input-block">
                <div id="rich_example">{$item['content_example']|raw}</div>
                <textarea style="display: none;" id="content_example" name="content_example">{$item['content_example']|raw}</textarea>
            </div>
        </div>		

       <div class="pop-bottom-bar">
            <button type="submit" class="layui-btn layui-btn-normal" lay-submit="" lay-filter="formSubmit" hisi-data="{pop: true, refresh: true}">提交保存</button>
            <a href="javascript:parent.layui.layer.closeAll();" class="layui-btn layui-btn-primary ml10">取消</a>
        </div>
    </div>
</form>
{include file="system@block/layui" /}
<script type="text/javascript">
    layui.use(['upload','layer'], function () {
        let layer = layui.layer,$ = layui.jquery, upload = layui.upload;
		initRichText()
        let uploadInst = upload.render({
            elem: '#img_btn'
            , exts: 'png|jpeg|jpg' 
            , url: "{:url('/system/file/put')}" 
            , done: function (res) {
                if (res.recode !== 0) {
                    return layer.msg('上传失败');
                }
                $('#img_show_example').attr('src', res.data.url);
                $('[name="img_example"]').val(res.data.url)
            }
            , error: function () {
                //演示失败状态，并实现重传
                layer.msg('添加失败！');
            }
        });
    });
</script>

<!--富文本绑定的代码 示例 无用则删除-->
<script src="__PUBLIC_JS__/editor/wangEditor/wangEditor.js"></script>
<script type="text/javascript">
    // 初始化富文本
    function initRichText() {
        const E = window.wangEditor
        const editor = new E('#rich_example')
        let $ = layui.jquery;
        // 自定义上传文件
        editor.config.customUploadImg = function (resultFiles, insertImgFn) {
            for (let index in resultFiles) {
                let formData = new FormData();                      // 创建一个form类型的数据
                formData.append('file',resultFiles[index]);
                $.ajax({
                    type: "POST",
                    url: "{:url('/system/file/put')}",
                    processData: false, // 将数据转换成对象，不对数据做处理，故 processData: false
                    contentType: false,   // 不设置数据类型
                    xhrFields: {                // 这样在请求的时候会自动将浏览器中的cookie发送给后台
                        withCredentials: true
                    },
                    data: formData,
                    success: function (res) {
                        insertImgFn(res.data.url)
                    }, error: function (data) {
                        layer.msg("网络错误", {time: 1500});
                    }
                })
            }
        }
		let content = $('#content_example');
        editor.config.onchange = function (html) {
            content.val(html)
        }
        editor.create()
        editor.txt.html(content.val())
    }
</script>
`

const Create = `<form class="layui-form alignment layui-form-pane" action="{:url('create')}" method="post">
    <div class="form_update">
		{$create}
		<!-- 上传图片的示例 -->
		<div class="layui-form-item">
            <label class="layui-form-label">上传图片示例</label>
            <div class="layui-input-block">
                <button type="button" class="layui-btn" id="img_btn">请选择头像</button>
				 <div class="ssy_img_parent">
						<img  width="40" class="ssy_img" height="40" id="img_show_example"/>
				</div>
            </div>
			<!-- 上传的图片的字段 -->
            <input type="hidden"  name="img_example" />
        </div>

		<!-- 单选的示例 -->
		<div class="layui-form-item">
            <label class="layui-form-label">单选的示例</label>
            <div class="layui-input-block">
                {foreach $radio_example as $key=>$vo }
                <input type="radio" name="radio_example" value="{$key}" title="{$vo}">
                {/foreach}
            </div>
        </div>
		
		<!-- 下拉选择的示例 -->
		<div class="layui-form-item">
            <label class="layui-form-label">下拉选择的示例</label>
            <div class="layui-input-block">
				<select name="select_example">
 				    {foreach $select_example as $key=>$vo }
					<option value="{$key}">{$vo}</option>
					{/foreach}
				</select>
            </div>
        </div>

		<!-- 富文本的例子 -->		
		 <div class="layui-form-item">
            <label class="layui-form-label">富文本的例子</label>
            <div class="layui-input-block">
                <div id="rich_example">{$content_example|raw}</div>
                <textarea style="display: none;" id="content_example" name="content_example">{$content_example}</textarea>
            </div>
        </div>		

		<div class="pop-bottom-bar">
            <button type="submit" class="layui-btn layui-btn-normal" lay-submit="" lay-filter="formSubmit" hisi-data="{pop: true, refresh: true}">添加</button>
            <a href="javascript:parent.layui.layer.closeAll();" class="layui-btn layui-btn-primary ml10">取消</a>
        </div>
    </div>
</form>
{include file="system@block/layui" /}
<script type="text/javascript">
    layui.use(['upload','layer'], function () {
        let layer = layui.layer,$ = layui.jquery, upload = layui.upload;
		initRichText()
        let uploadInst = upload.render({
            elem: '#img_btn'
            , exts: 'png|jpeg|jpg' 
            , url: "{:url('/system/file/put')}" 
            , done: function (res) {
                if (res.recode !== 0) {
                    return layer.msg('上传失败');
                }
                $('#img_show_example').attr('src', res.data.url);
                $('[name="img_example"]').val(res.data.url)
            }
            , error: function () {
                //演示失败状态，并实现重传
                layer.msg('添加失败！');
            }
        });
    });
</script>

<!--富文本绑定的代码 示例 无用则删除-->
<script src="__PUBLIC_JS__/editor/wangEditor/wangEditor.js"></script>
<script type="text/javascript">
    // 初始化富文本
    function initRichText() {
        const E = window.wangEditor
        const editor = new E('#rich_example')
        let $ = layui.jquery;
        // 自定义上传文件
        editor.config.customUploadImg = function (resultFiles, insertImgFn) {
            for (let index in resultFiles) {
                let formData = new FormData();                      // 创建一个form类型的数据
                formData.append('file',resultFiles[index]);
                $.ajax({
                    type: "POST",
                    url: "{:url('/system/file/put')}",
                    processData: false, // 将数据转换成对象，不对数据做处理，故 processData: false
                    contentType: false,   // 不设置数据类型
                    xhrFields: {                // 这样在请求的时候会自动将浏览器中的cookie发送给后台
                        withCredentials: true
                    },
                    data: formData,
                    success: function (res) {
                        insertImgFn(res.data.url)
                    }, error: function (data) {
                        layer.msg("网络错误", {time: 1500});
                    }
                })
            }
        }
		let content = $('#content_example');
        editor.config.onchange = function (html) {
            content.val(html)
        }
        editor.create()
        editor.txt.html(content.val())
    }
</script>`
