package make

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"yii/make/laravel/util"
	util_tp "yii/make/tp/util"
)

// 默认生成在storage/apizzat目录下
type Apizzat struct {
	Name      string
	Namespace string
	Root      string
	DbPrefix  string

	fieldsMap map[string]*util.Field
	fields    []*util.Field
	Comment   string
}

func (m *Apizzat) Make() {
	//m.Root = "C:\\Users\\yinhe\\Desktop\\projects\\ServerSiteMirocs"
	tableName := util_tp.Camel2Case(m.Name, '_')
	db, error := util.MysqlConnect(util.LoadConfigEnvLaravel(m.Root, m.DbPrefix))
	if error != nil {
		log.Fatal(error.Error())
		os.Exit(1)
	}
	m.fieldsMap, m.fields, m.Comment = util.TableFieldsMap(db, tableName)
	var content []byte
	if m.Namespace == "request" {
		content = m.request()
	} else {
		content = m.response()
	}
	filepath := util.GetFilePath(m.Root, "storage", "apizzat", m.Namespace, tableName)
	if !util_tp.WriteFile(filepath, string(content)) {
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
