package tp_make

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	tp_template "yii/make/tp/template"
	util_tp "yii/make/tp/util"
)

type Model struct {
	Name         string
	NameSpace    string
	NameSpaces   []string
	Table        string
	Root         string
	Reset        bool
	mp           map[string]*util_tp.Field
	fields       []*util_tp.Field
	tableComment string
}

func (m *Model) Make() {
	tableName := util_tp.Camel2Case(m.Table, '_')
	m.mp, m.fields, m.tableComment = util_tp.TableFieldsMap(m.Root, tableName, m.NameSpaces)
	filepath := util_tp.GetFilePath(m.Root, m.NameSpace, m.Name+".php")
	property := m.property()
	content := fmt.Sprintf(tp_template.Model, m.NameSpace, property, m.Name)
	if !util_tp.WriteFile(filepath, content) && property != "" {
		re := regexp.MustCompile(fmt.Sprintf("/\\*\\*[^`]*?class %s", m.Name))
		ori, _ := ioutil.ReadFile(filepath)
		dest := strings.ReplaceAll(string(ori), re.FindString(string(ori)), fmt.Sprintf("%s\nclass %s", property, m.Name))
		if dest != string(ori) {
			_ = ioutil.WriteFile(filepath, []byte(dest), 0777)
		}
	}
}

func (m Model) property() string {
	str := fmt.Sprintf("/**\n * %s\n", m.tableComment)
	for i := 0; i < len(m.fields); i++ {
		str = str + fmt.Sprintf(" * @property %s $%s %s\n", m.fields[i].DataType, m.fields[i].ColumnName, m.fields[i].ColumnComment)
	}
	if str == "/**\n" {
		return ""
	}
	return str + " */"
}
