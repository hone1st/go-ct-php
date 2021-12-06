package make

import (
	"fmt"
	"honelst/make/gii/template"

	"honelst/util/mysql"
	"honelst/util/other"
	"strings"
)

type AdminModel struct {
	Model        string
	Module       string
	Mp           map[string]*mysql.Field
	Fields       []*mysql.Field
	Path         string
	TableComment string
}

func (m AdminModel) Make() {
	property := ""
	if len(m.Mp) > 0 {
		property = m.property(m.TableComment)
	}
	fileContent := template.Model
	fileContent = strings.ReplaceAll(fileContent, "{$model}", m.Model)
	fileContent = strings.ReplaceAll(fileContent, "{$module}", m.Module)
	fileContent = strings.ReplaceAll(fileContent, "{$property}", property)
	other.WriteFile(m.Path, fileContent)
}

func (m AdminModel) property(tableComment string) string {
	str := fmt.Sprintf("/**\n * %s\n", tableComment)
	for i := 0; i < len(m.Fields); i++ {
		str = str + fmt.Sprintf(" * @property %s $%s %s\n", m.Fields[i].DataType, m.Fields[i].ColumnName, m.Fields[i].ColumnComment)
	}
	if str == "/**\n" {
		return ""
	}
	return str + " */"
}
