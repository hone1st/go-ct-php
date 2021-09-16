package make

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"yii/make/laravel/template"
	"yii/make/laravel/util"
	util_tp "yii/make/tp/util"
)

type Model struct {
	Name       string
	Namespace  string
	Root       string
	Table      string
	NameSpaces []string
	DbPrefix   string

	fieldsMap map[string]*util.Field
	fields    []*util.Field
	Comment   string
}

func (m *Model) Make() {
	tableName := util_tp.Camel2Case(m.Name, '_')
	db, error := util.MysqlConnect(util.LoadConfigEnvLaravel(m.Root, m.DbPrefix))
	if error != nil {
		log.Fatal(error.Error())
		os.Exit(1)
	}
	m.fieldsMap, m.fields, m.Comment = util.TableFieldsMap(db, tableName)
	rules := m.fieldsRule()
	casts := m.fieldsTrans()
	fillable := m.fillable()
	property := m.property()
	filepath := util.GetFilePath(m.Root, m.Namespace, m.Name+".php")

	content := strings.ReplaceAll(template.Model, "{$namespace}", m.Namespace)
	content = strings.ReplaceAll(content, "{$model}", m.Name)
	content = strings.ReplaceAll(content, "{$property}", property)
	content = strings.ReplaceAll(content, "{$table}", tableName)
	content = strings.ReplaceAll(content, "{$fields}", fillable)
	content = strings.ReplaceAll(content, "{$fieldsTrans}", casts)
	content = strings.ReplaceAll(content, "{$fieldsRule}", rules)
	if !util_tp.WriteFile(filepath, content) && property != "" {
		re := regexp.MustCompile(fmt.Sprintf("/\\*\\*[^`]*?class %s", m.Name))
		ori, _ := ioutil.ReadFile(filepath)
		dest := strings.ReplaceAll(string(ori), re.FindString(string(ori)), fmt.Sprintf("%s\nclass %s", property, m.Name))
		if dest != string(ori) {
			_ = ioutil.WriteFile(filepath, []byte(dest), 0777)
		}
	}

}

func (m *Model) property() string {
	str := fmt.Sprintf("/**\n * %s\n", m.Comment)
	for i := 0; i < len(m.fields); i++ {
		field := m.fields[i]
		field.ColumnComment = strings.ReplaceAll(field.ColumnComment, "\r\n", "")
		if m.fields[i].DataType == "datetime" {
			str = str + fmt.Sprintf(" * @property %s $%s %s\n", "string", field.ColumnName, field.ColumnComment)
		} else {
			str = str + fmt.Sprintf(" * @property %s $%s %s\n", field.DataType, field.ColumnName, field.ColumnComment)
		}
	}
	if str == "/**\n" {
		return ""
	}
	return str + " */"
}

func (m *Model) fillable() string {
	fields := make([]string, 0)
	for field, _ := range m.fieldsMap {
		fields = append(fields, fmt.Sprintf("        '%s',", field))
	}
	return strings.Join(fields, "\r\n")
}

func (m *Model) fieldsRule() string {
	str := `
		'deleted_at' => 'nullable',
		'created_at' => 'nullable',
		'updated_at' => 'nullable',
`
	for field, v := range m.fieldsMap {
		if field == "created_at" || field == "updated_at" || field == "deleted_at" {
			continue
		}
		rules := make([]string, 0)
		if v.DataType == "string" && v.CharacterMaximumLength > 0 {
			rules = append(rules, "required", v.DataType)
			rules = append(rules, fmt.Sprintf("max:%d", v.CharacterMaximumLength))
		} else if v.DataType == "int" {
			rules = append(rules, "required", "integer")
		} else {
			rules = append(rules, "required", v.DataType)
		}
		str = str + fmt.Sprintf("\t\t'%s' => '%s',\r\n", field, strings.Join(rules, "|"))
	}
	return str
}

func (m *Model) fieldsTrans() string {
	casts := make([]string, 0)
	for field, v := range m.fieldsMap {
		if v.DataType == "int" {
			casts = append(casts, fmt.Sprintf("        '%s' => '%s',", field, "integer"))
		} else {
			casts = append(casts, fmt.Sprintf("        '%s' => '%s',", field, v.DataType))
		}
	}

	return strings.Join(casts, "\r\n")
}
