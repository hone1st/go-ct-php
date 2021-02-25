```TXT
 * gii model:portal\modules\v2\models@ApiTest 生成model
 * gii controller:portal\modules\v2\controllers@ApiTest 生成controller
 * gii service:portal\modules\v2\controllers\service@ApiTest 生成service
 * gii form:portal\modules\v2\form@ApiTest 生成form
 * gii module:portal\modules\v2\modules@test 自动生成portal\modules\v2\modules\test\Module.php 
```


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