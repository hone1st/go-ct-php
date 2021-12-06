package make

import (
	"flag"
	"log"
	"strings"
)

var (
	//db string
	g string
)

// model:portal\modules\v2\models@ApiTest   					规则model文件存在的命名空间和model名字
// controller:portal\modules\v2\controllers@ApiTest				规则controller文件存在的命名空间和controller名字
// service:portal\modules\v2\controllers\service@ApiTest		规则service文件存在的命名空间和service名字
// form:portal\modules\v2\form@ApiTest							规则form文件存在的命名空间和form名字

// 有点绕 不推荐一次性生成所有 暂注释了module生成所有basic文件
// all:portal\modules\v2@ApiTest								规则生成model/controller/service/form文件 portal\modules\v2表示某个模块的命名空间
// module:portal\modules\v2\modules@test						规则生成model/controller/service/form文件 portal\modules\v2\modules表示某个父类模块的子模块的目录

func argsParse() {
	//flag.StringVar(&db, "db", "", "指定数据库链接，默认是env的默认链接")
	flag.StringVar(&g, "g", "", `执行的参数 
-g 
// model:portal\modules\v2\models@ApiTest   					规则model文件存在的命名空间和model名字
// controller:portal\modules\v2\controllers@ApiTest				规则controller文件存在的命名空间和controller名字
// service:portal\modules\v2\controllers\service@ApiTest		规则service文件存在的命名空间和service名字
// form:portal\modules\v2\form@ApiTest							规则form文件存在的命名空间和form名字

// 有点绕 不推荐一次性生成所有 暂注释了module生成所有basic文件
// all:portal\modules\v2@ApiTest								规则生成model/controller/service/form文件 portal\modules\v2表示某个模块的命名空间
// module:portal\modules\v2\modules@test						规则生成model/controller/service/form文件 portal\modules\v2\modules表示某个父类模块的子模块的目录

`)
	flag.Parse()
}

func Make(root string) {
	argsParse()
	if !strings.Contains(g, ":") && !strings.Contains(g, "@") {
		log.Fatalln("-g 的值格式非法")
	}
	action := strings.Split(g, ":")
	if len(action) != 2 {
		log.Fatalln("-g 的值格式非法")
	}
	dos := strings.Split(action[1], "@")
	if len(dos) != 2 {
		log.Fatalln("-g 的值格式非法")
	}

	namespaces := strings.Split(dos[0], "\\")
	if len(namespaces) > 0 && namespaces[0] == "" {
		namespaces = namespaces[1:]
	}
	namespace := strings.Join(namespaces, "\\")
	name := dos[1]
	// repository|validation|controller|model
	// dos[0] 是命名空间 dos[1] 是文件名字或者表名字

	// apizzat
	// dos[0] 生成request还是response dos[1] 表名字
	do(root, action[0], namespace, name)
}

func do(root string, action string, namespace string, name string) {
	namespaces := strings.Split(namespace, "\\")
	if len(namespaces) > 0 && namespaces[0] == "" {
		namespaces = namespaces[1:]
	}
	namespace = strings.Join(namespaces, "\\")
	switch strings.ToLower(action) {
	case "model":
		(&Model{
			Name:      name,
			NameSpace: namespace,
			Root:      root,
		}).Make()
		break
	case "reset":
		(&Controller{
			Name:       name,
			NameSpace:  namespace,
			Root:       root,
			NameSpaces: namespaces,
			Reset:      true,
		}).Make()
		break
	case "controller":
		(&Controller{
			Name:       name,
			NameSpace:  namespace,
			Root:       root,
			NameSpaces: namespaces,
		}).Make()
		break
	case "service":
		(&Service{
			Name:      name,
			NameSpace: namespace,
			Root:      root,
		}).Make()
		break
	case "form":
		(&Form{
			Name:      name,
			NameSpace: namespace,
			Root:      root,
		}).Make()
		break
	case "module":
		// doBasic(fmt.Sprintf("%s\\%s", namespace, name), name)
		(&Module{
			Name:       name,
			NameSpace:  namespace,
			Root:       root,
			NameSpaces: namespaces,
		}).Make()
		break
	case "all":
		doBasic(root, namespace, name)
		break
	default:
		log.Fatalf("非法操作：%s \r\n", action[0])
	}
}

func doBasic(root, namespace, name string) {
	do(root, "model", namespace+"\\models", name)
	do(root, "controller", namespace+"\\controllers", name)
	do(root, "service", namespace+"\\controllers\\services", name)
	do(root, "form", namespace+"\\forms", name)
}
