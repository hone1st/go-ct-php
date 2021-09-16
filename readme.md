```
go build 生成可执行文件
编译window可执行文件
set GOARCH=386
set GOOS=windows

编译linux可执行文件
SET GOOS=linux
SET GOARCH=amd64

编译mac可执行文件
SET GOOS=darwin
SET GOARCH=amd64
```

```gii 
yii2框架的代码生成
# cd /make/gii 
# go build . 
生成可执行文件
```
```TXT
 * gii model:portal\modules\v2\models@ApiTest 生成model
 * gii controller:portal\modules\v2\controllers@ApiTest 生成controller
 * gii service:portal\modules\v2\controllers\service@ApiTest 生成service
 * gii form:portal\modules\v2\form@ApiTest 生成form
 * gii module:portal\modules\v2\modules@test 自动生成portal\modules\v2\modules\test\Module.php 
```


```
tp5.1框架生成代码
# cd /make/tp 
# go build . 

tp5.1配合hisiphpv2快速生成curd代码 
module表示模块名字
model表示表明去掉数据库配置文件设置的前缀的驼峰
## tp.exe curd:app\module@model
例如：在user模块下生成user表的curd
[
    prifix => 'prefix' // 配置文件
]
prfix_user => User // 真实表面

## tp.exe curd:app\user@User

tp5.1所有通用的model文件生成
## tp.exe model:namespace@model
namespace 表示要生成的模型文件的命名空间
model表示表名字

例如生成user表
## tp.exe model:app\user\model@User

```


```laravel
laravel生成的命令

main.exe 命令大全

-g 是生成文件的东西
							   生成模型 命名空间  表名驼峰	
	生成模型的例子: main.exe -g model:App\Models@Order 

	生成apizzat的(request/response)字典  main.exe -g apizzat:request/response@Order

-db 指定env的数据库的前缀

	例如指定数据库生成模型
									   生成模型 命名空间  表名驼峰	
	生成模型的例子: main.exe -g model:App\Models@User  -db CRM_
	生成apizzat的(request/response)字典  main.exe -g apizzat:request/response@Order -db CRM_
	

```