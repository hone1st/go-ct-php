package make

import (
	"fmt"
	"strings"
	"yii/make/laravel/template"
	"yii/make/laravel/util"
	util_tp "yii/make/tp/util"
)

type Repository struct {
	Name       string
	Namespace  string
	Root       string
	Table      string
	NameSpaces []string
	DbPrefix   string
}

func (m *Repository) Make() {
	filepath := util.GetFilePath(m.Root, m.Namespace, m.Name+"Repository.php")
	content := strings.ReplaceAll(template.Repository, "{$namespace}", m.Namespace)
	content = strings.ReplaceAll(content, "{$name}", m.Name)
	tableName := util_tp.Camel2Case(m.Name, '_')
	content = strings.ReplaceAll(content, "{$name_id}", tableName+"_id")
	// 检测一下是否存在这张表
	db, error := util.MysqlConnect(util.LoadConfigEnvLaravel(m.Root, m.DbPrefix))
	fieldsMap := ""
	if error == nil {
		_, fields, _ := util.TableFieldsMap(db, tableName)
		if len(fields) > 0 {
			line := make([]string, 0)
			for i := 0; i < len(fields); i++ {
				field := fields[i]
				if field.ColumnName == "id" {
					continue
				}
				if field.ColumnName == "created_at" {
					continue
				}
				if field.ColumnName == "updated_at" {
					continue
				}
				if field.ColumnName == "deleted_at" {
					continue
				}
				line = append(line, fmt.Sprintf("            '%s'          => $input['%s'],", field.ColumnName, field.ColumnName))
			}
			fieldsMap = strings.Join(line, "\r\n")
		}
	}
	content = strings.ReplaceAll(content, "{$fields_map}", fieldsMap)
	util_tp.WriteFile(filepath, content)
}
