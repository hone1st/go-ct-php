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

func argsParse() {
	//flag.StringVar(&db, "db", "", "指定数据库链接，默认是env的默认链接")
	flag.StringVar(&g, "g", "", `执行的参数 model/apizzat:namespace/(request/response)@table
-g model:App\Models@User  生成模型
-g controller:App\Http\Controllers@User  生成控制器
-g repository:App\Repository@User  生成逻辑层
-g apizzat:response@User  生成对象定义
-g validation:App\Http\Controllers\Backend@Question  指定某个控制器的文件生成对于的验证层的代码
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
	switch action[0] {
	case "model":
		(&Model{
			Name:       name,
			NameSpace:  namespace,
			Root:       root,
			NameSpaces: namespaces,
		}).Make()

		break
	case "curd":
		(&Curd{
			Name:       name,
			NameSpace:  namespace,
			Root:       root,
			NameSpaces: namespaces,
		}).Make()
		break
	default:
		log.Fatalf("非法操作：%s \r\n", action[0])
	}
}
