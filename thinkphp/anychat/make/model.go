package make

import (
	"fmt"
	"honelst/thinkphp/anychat/template"
	"honelst/util/mysql"
	"honelst/util/other"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type Model struct {
	Name         string
	NameSpace    string
	NameSpaces   []string
	Table        string
	Root         string
	Reset        bool
	mp           map[string]*mysql.Field
	fields       []*mysql.Field
	tableComment string
}

func (m *Model) Make() {
	tableName := other.Camel2Case(m.Table, '_')
	db, err := mysql.MysqlConnect(mysql.LoadConfigEnvTp(m.Root))
	if err != nil {
		log.Fatal(err.Error())
	}
	m.mp, m.fields, m.tableComment = mysql.TableFieldsMap(db, tableName)
	filepath := other.GetFilePath(m.Root, m.NameSpace, m.Name+".php")
	property := m.property()
	content := fmt.Sprintf(template.Model, m.NameSpace, property, m.Name)
	if !other.WriteFile(filepath, content) && property != "" {
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
