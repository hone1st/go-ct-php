package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"yii/make/inter"
	"yii/make/tp/admin_make"
	tp_make "yii/make/tp/make"
)

var (
	root string
	make inter.Make
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("必须有一个参数!")
		os.Exit(0)
	}
	root = path.Dir(args[0])
	first := args[1]
	two := ""
	if len(args) == 3 {
		two = args[2]
	}
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
	do(whats[0], dos[0], dos[1], two)
}

func do(what, namespace, name, two string) {
	if two == "" {
		two = name
	}
	namespaces := strings.Split(namespace, "\\")
	if len(namespaces) > 0 && namespaces[0] == "" {
		namespaces = namespaces[1:]
	}
	namespace = strings.Join(namespaces, "\\")

	switch strings.ToLower(what) {
	case "model":
		make = &tp_make.Model{
			Name:       name,
			NameSpace:  namespace,
			Root:       root,
			Table:      two,
			NameSpaces: namespaces,
		}
		break
	case "curd":
		make = &admin_make.Curd{
			Name:       name,
			NameSpace:  namespace,
			Root:       root,
			Table:      two,
			NameSpaces: namespaces,
		}
		break
	}
	make.Make()
}
