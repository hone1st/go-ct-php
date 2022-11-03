package make

import (
	"fmt"
	"honelst/laravel/galaxy_new/template"
	"honelst/util/mysql"
	"honelst/util/other"
	"strings"
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
	filepath := other.GetFilePath(m.Root, m.Namespace, m.Name+"Repository.php")
	content := strings.ReplaceAll(template.Repository, "{$namespace}", other.CapOrLow(m.Namespace, true))
	content = strings.ReplaceAll(content, "{$name}", m.Name)
	tableName := other.Camel2Case(m.Name, '_')
	content = strings.ReplaceAll(content, "{$name_id}", tableName+"_id")
	// 检测一下是否存在这张表
	db, error := mysql.MysqlConnect(mysql.LoadConfigEnvLaravel(m.Root, m.DbPrefix))
	fieldsMap := ""
	when := ""
	if error == nil {
		_, fields, _ := mysql.TableFieldsMap(db, tableName)
		if len(fields) > 0 {
			whenLine := make([]string, 0)
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
				whenLine = append(whenLine, fmt.Sprintf("            ->when(isset($params['%s']) && !empty($params['%s']), function ($query) use ($params) {\n                $query->where('%s', $params['%s']);\n            })", field.ColumnName, field.ColumnName, field.ColumnName, field.ColumnName))

			}
			fieldsMap = strings.Join(line, "\r\n")
			when = strings.Join(whenLine, "\r\n")
		}
	}
	content = strings.ReplaceAll(content, "{$fields_map}", fieldsMap)
	content = strings.ReplaceAll(content, "{$when}", when)
	other.WriteFile(filepath, content)
}
