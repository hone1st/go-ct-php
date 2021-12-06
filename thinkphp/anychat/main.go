package main

import (
	"honelst/thinkphp/anychat/make"
	"os"
	"path"
)

// 就职-suishiliao

// 初衷：
// 		1.减少cv重复的代码编写
// 		2.提高开发效率
// 		3.减少加班
// 		4.多点时间学习其他的知识

func main() {
	root := path.Dir(os.Args[0])
	make.Make(root)
}
