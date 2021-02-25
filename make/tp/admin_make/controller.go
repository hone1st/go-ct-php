package admin_make

import (
	"strings"
	"yii/make/tp/admin_template"
	util_tp "yii/make/tp/util"
)

type Controller struct {
	LModel string
	Model  string
	Module string
	Path   string
}

func (c Controller) Make() {
	// 替换模板的内容
	fileContent := admin_template.Controller
	fileContent = strings.ReplaceAll(fileContent, "{$lModel}", c.LModel)
	fileContent = strings.ReplaceAll(fileContent, "{$model}", c.Model)
	fileContent = strings.ReplaceAll(fileContent, "{$module}", c.Module)
	util_tp.WriteFile(c.Path, fileContent)
}
