package make

import (
	"encoding/json"
	"fmt"
	"honelst/util/mysql"
	"honelst/util/other"
	"io/ioutil"
	"log"
	"strings"
)

// 默认生成在storage/apizzat目录下

type Apizzat struct {
	Name      string
	Namespace string
	Root      string
	DbPrefix  string

	fieldsMap map[string]*mysql.Field
	fields    []*mysql.Field
	Comment   string
}

func (m *Apizzat) Make() {
	tableName := other.Camel2Case(m.Name, '_')
	db, error := mysql.MysqlConnect(mysql.LoadConfigEnvLaravel(m.Root, m.DbPrefix))
	if error != nil {
		log.Fatal(error.Error())
	}
	m.fieldsMap, m.fields, m.Comment = mysql.TableFieldsMap(db, tableName)
	var content []byte
	if m.Namespace == "request" {
		content = m.request()
	} else {
		content = m.response()
	}
	filepath := other.GetFilePath(m.Root, "storage", "apizzat", m.Namespace, tableName)
	if !other.WriteFile(filepath, string(content)) {
		_ = ioutil.WriteFile(filepath, content, 0777)
		fmt.Printf("文件：【%s】 已重新生成成功！\n", filepath)
	}
}

type Request struct {
	Key   string        `json:"key"`
	Desc  string        `json:"desc"`
	Eg    string        `json:"eg"`
	Enums []interface{} `json:"enums"`
	Type  string        `json:"type"`
}

func (m Apizzat) request() []byte {
	mp := make([]*Request, 0)
	for i := 0; i < len(m.fields); i++ {
		field := m.fields[i]
		mp = append(mp, &Request{
			Key:   field.ColumnName,
			Desc:  field.ColumnComment,
			Eg:    "",
			Enums: nil,
			Type:  field.DataType,
		})
	}

	marshal, _ := json.MarshalIndent(mp, "\t", "\t")
	return marshal
}

// ex field:comment:type
func (m *Apizzat) response() []byte {
	strArr := make([]string, 0)
	for i := 0; i < len(m.fields); i++ {
		field := m.fields[i]
		field.ColumnComment = strings.ReplaceAll(field.ColumnComment, "\r\n", "")
		strArr = append(strArr, strings.Join([]string{field.ColumnName, field.ColumnComment, field.DataType}, ":"))
	}
	return []byte(strings.Join(strArr, "\r\n"))
}
