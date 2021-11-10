package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"yii/make/inter"
	make2 "yii/make/laravel/make"
)

var (
	root string
	make inter.Make
	db   string
	g    string
)

func argsParse() {
	flag.StringVar(&db, "db", "", "指定数据库链接，默认是env的默认链接")
	flag.StringVar(&g, "g", "", `执行的参数 model/apizzat:namespace/(request/reponse)@table
-g model:App\Models@User  生成模型
-g controller:App\Http\Controllers@User  生成控制器
-g repository:App\Repository@User  生成逻辑层
-g apizzat:response@User  生成对象定义
`)
	flag.Parse()
}

func main() {
	argsParse()
	root = path.Dir(os.Args[0])
	if !strings.Contains(g, ":") && !strings.Contains(g, "@") {
		fmt.Println("参数非法!")
		os.Exit(0)
	}
	whats := strings.Split(g, ":")
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
			Name:       name,
			Namespace:  namespace,
			Root:       root,
			NameSpaces: namespaces,
			DbPrefix:   db,
		}
		break
	case "apizzat":
		make = &make2.Apizzat{
			Name:      name,
			Namespace: namespace,
			Root:      root,
			DbPrefix:  db,
		}
		break
	case "controller":
		make = &make2.Controller{
			Name:       name,
			Namespace:  namespace,
			Root:       root,
			DbPrefix:   db,
			NameSpaces: namespaces,
		}
		break
	case "repository":
		make = &make2.Repository{
			Name:       name,
			Namespace:  namespace,
			Root:       root,
			DbPrefix:   db,
			NameSpaces: namespaces,
		}
		break
	default:
		fmt.Println("暂不支持!")
		os.Exit(0)
	}
	make.Make()
}
