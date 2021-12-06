package make

import (
	"honelst/thinkphp/anychat/template"
	"honelst/util/other"
	"strings"
)

type AdminService struct {
	Model  string
	Module string
	Path   string
}

func (s AdminService) Make() {
	fileContent := template.AdminService
	fileContent = strings.ReplaceAll(fileContent, "{$model}", s.Model)
	fileContent = strings.ReplaceAll(fileContent, "{$module}", s.Module)
	other.WriteFile(s.Path, fileContent)
}
