package make

import (
	"fmt"
	"honelst/thinkphp/anychat/template"
	"honelst/util/mysql"
	"honelst/util/other"
	"strings"
)

type AdminValidate struct {
	Model  string
	Module string
	Mp     map[string]*mysql.Field
	Fields []*mysql.Field
	Path   string
}

func (v AdminValidate) Make() {
	// 替换模板的内容
	fileContent := template.AdminValidate
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
	other.WriteFile(v.Path, fileContent)
}
