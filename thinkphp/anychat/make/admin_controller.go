package make

import (
	"honelst/thinkphp/anychat/template"
	"honelst/util/other"
	"strings"
)

type AdminController struct {
	LModel string
	Model  string
	Module string
	Path   string
}

func (c AdminController) Make() {
	// 替换模板的内容
	fileContent := template.AdminController
	fileContent = strings.ReplaceAll(fileContent, "{$lModel}", c.LModel)
	fileContent = strings.ReplaceAll(fileContent, "{$model}", c.Model)
	fileContent = strings.ReplaceAll(fileContent, "{$module}", c.Module)
	other.WriteFile(c.Path, fileContent)
}
