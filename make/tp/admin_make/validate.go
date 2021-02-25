package admin_make

import (
	"fmt"
	"strings"
	"yii/make/tp/admin_template"
	util_tp "yii/make/tp/util"
)

type Validate struct {
	Model  string
	Module string
	Mp     map[string]*util_tp.Field
	Fields []*util_tp.Field
	Path   string
}

func (v Validate) Make() {
	// 替换模板的内容
	fileContent := admin_template.Validate
	fileContent = strings.ReplaceAll(fileContent, "{$model}", v.Model)
	fileContent = strings.ReplaceAll(fileContent, "{$module}", v.Module)
	if len(v.Mp) > 0 {
		rule := make([]string, 0)
		message := make([]string, 0)
		scene := make([]string, 0)
		for name, field := range v.Mp {
			message = append(message, fmt.Sprintf("\t\t\t'%s' => '%s'", name, field.ColumnComment))
			scene = append(scene, fmt.Sprintf(" '%s'", name))
			rule = append(rule, fmt.Sprintf("\t\t\t'%s'  => []", name))
		}
		fileContent = strings.ReplaceAll(fileContent, "{$rule}", strings.Join(rule, ",\r\n"))
		fileContent = strings.ReplaceAll(fileContent, "{$scene}", strings.Join(scene, ","))
		fileContent = strings.ReplaceAll(fileContent, "{$message}", strings.Join(message, ",\r\n"))
	}
	util_tp.WriteFile(v.Path, fileContent)
}
