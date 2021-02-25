package admin_make

import (
	"fmt"
	"strings"
	"yii/make/tp/admin_template"
	util_tp "yii/make/tp/util"
)

type Model struct {
	Model        string
	Module       string
	Mp           map[string]*util_tp.Field
	Fields       []*util_tp.Field
	Path         string
	TableComment string
}

func (m Model) Make() {
	property := ""
	if len(m.Mp) > 0 {
		property = m.property(m.TableComment)
	}
	fileContent := admin_template.Model
	fileContent = strings.ReplaceAll(fileContent, "{$model}", m.Model)
	fileContent = strings.ReplaceAll(fileContent, "{$module}", m.Module)
	fileContent = strings.ReplaceAll(fileContent, "{$property}", property)
	util_tp.WriteFile(m.Path, fileContent)
}

func (m Model) property(tableComment string) string {
	str := fmt.Sprintf("/**\n * %s\n", tableComment)
	for i := 0; i < len(m.Fields); i++ {
		str = str + fmt.Sprintf(" * @property %s $%s %s\n", m.Fields[i].DataType, m.Fields[i].ColumnName, m.Fields[i].ColumnComment)
	}
	if str == "/**\n" {
		return ""
	}
	return str + " */"
}
