package admin_make

import (
	"strings"
	"yii/make/tp/admin_template"
	util_tp "yii/make/tp/util"
)

type Service struct {
	Model  string
	Module string
	Path   string
}

func (s Service) Make() {
	fileContent := admin_template.Service
	fileContent = strings.ReplaceAll(fileContent, "{$model}", s.Model)
	fileContent = strings.ReplaceAll(fileContent, "{$module}", s.Module)
	util_tp.WriteFile(s.Path, fileContent)
}
