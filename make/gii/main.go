/**
 * @Time : 2020/10/9 16:32
 * @Author : liang
 * @File : main
 * @Software: GoLand
 */

package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	make2 "yii/make/gii/make"
	"yii/make/inter"
)

var (
	root string
	make inter.Make
)

// model:portal\modules\v2\models@ApiTest   					规则model文件存在的命名空间和model名字
// controller:portal\modules\v2\controllers@ApiTest				规则controller文件存在的命名空间和controller名字
// service:portal\modules\v2\controllers\service@ApiTest		规则service文件存在的命名空间和service名字
// form:portal\modules\v2\form@ApiTest							规则form文件存在的命名空间和form名字

// 有点绕 不推荐一次性生成所有 暂注释了module生成所有basic文件
// all:portal\modules\v2@ApiTest								规则生成model/controller/service/form文件 portal\modules\v2表示某个模块的命名空间
// module:portal\modules\v2\modules@test						规则生成model/controller/service/form文件 portal\modules\v2\modules表示某个父类模块的子模块的目录

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("必须有一个参数!")
		os.Exit(0)
	}
	root = path.Dir(args[0])
	// root = "F:\\project\\ai_service_webv2_backend"
	first := args[1]
	if !strings.Contains(first, ":") && !strings.Contains(first, "@") {
		fmt.Println("参数非法!")
		os.Exit(0)
	}
	whats := strings.Split(first, ":")
	if len(whats) != 2 {
		fmt.Println("参数非法!")
		os.Exit(0)
	}
	dos := strings.Split(whats[1], "@")
	if len(dos) != 2 {
		fmt.Println("参数非法!")
		os.Exit(0)
	}
	do(whats[0], dos[0], dos[1])
}

func do(what, namespace, name string) {

	namespaces := strings.Split(namespace, "\\")
	if len(namespaces) > 0 && namespaces[0] == "" {
		namespaces = namespaces[1:]
	}
	namespace = strings.Join(namespaces, "\\")

	switch strings.ToLower(what) {
	case "model":
		make = &make2.Model{
			Name:      name,
			NameSpace: namespace,
			Root:      root,
		}
		break
	case "reset":
		make = &make2.Controller{
			Name:       name,
			NameSpace:  namespace,
			Root:       root,
			NameSpaces: namespaces,
			Reset:      true,
		}
		break
	case "controller":
		make = &make2.Controller{
			Name:       name,
			NameSpace:  namespace,
			Root:       root,
			NameSpaces: namespaces,
		}
		break
	case "service":
		make = &make2.Service{
			Name:      name,
			NameSpace: namespace,
			Root:      root,
		}
		break
	case "form":
		make = &make2.Form{
			Name:      name,
			NameSpace: namespace,
			Root:      root,
		}
		break
	case "module":
		// doBasic(fmt.Sprintf("%s\\%s", namespace, name), name)
		make = &make2.Module{
			Name:       name,
			NameSpace:  namespace,
			Root:       root,
			NameSpaces: namespaces,
		}
		break
	case "all":
		doBasic(namespace, name)
		break
	}
	make.Make()
}

func doBasic(namespace, name string) {
	do("model", namespace+"\\models", name)
	do("controller", namespace+"\\controllers", name)
	do("service", namespace+"\\controllers\\services", name)
	do("form", namespace+"\\forms", name)
}
