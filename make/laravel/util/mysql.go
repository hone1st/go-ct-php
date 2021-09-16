package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"unicode"
)

type Field struct {
	ColumnName             string `json:"column_name"`
	ColumnComment          string `json:"column_comment"`
	DataType               string `json:"data_type"`
	CharacterMaximumLength int    `json:"character_maximum_length"`
}

const Basic = "int,tinyint,bigint"

// 数据库转化为php的类型
func typeToPhp(t string) string {
	switch t {
	case "varchar":
	case "char":
		return "string"
	case "int":
	case "tinyint":
		return "int"
	case "decimal":
		return "float"
	case "datetime":
		return "datetime"
	}
	return "string"
}

func TableFieldsMap(db *DbResult, tableName string) (map[string]*Field, []*Field, string) {

	mp := make(map[string]*Field, 0)
	fields := make([]*Field, 0)
	if db == nil {
		return mp, fields, ""
	}
	res, err := db.Db.Query("select column_name,column_comment,data_type,character_maximum_length from "+
		"information_schema.columns where table_schema =? and table_name = ? ; ", db.Config.database, db.Config.prefix+tableName)
	if err != nil {
		log.Println(err)
	}
	tableRes, err := db.Db.Query("select table_comment from "+
		"information_schema.TABLES where table_schema =? and table_name = ? ; ", db.Config.database, db.Config.prefix+tableName)
	if err != nil {
		log.Println(err)
	}
	defer db.Db.Close()
	defer res.Close()
	defer tableRes.Close()

	var tableComment string
	for tableRes.Next() {
		tableRes.Scan(&tableComment)
	}
	var columnName, columnComment, dataType string
	var characterMaximumLength int
	for res.Next() {
		_ = res.Scan(&columnName, &columnComment, &dataType, &characterMaximumLength)
		f := &Field{
			ColumnName:             columnName,
			ColumnComment:          columnComment,
			DataType:               dataType,
			CharacterMaximumLength: characterMaximumLength,
		}
		f.DataType = typeToPhp(f.DataType)
		if _, ex := mp[f.ColumnName]; !ex {
			mp[f.ColumnName] = f
		}
		fields = append(fields, f)
	}
	return mp, fields, tableComment

}

// 获取相同的类型的值
func getValueByRe(reContent string) string {
	result := strings.Split(reContent, "=")
	dest := result[1]
	return strings.TrimSpace(dest)
}

// 生成文件
func WriteFile(filepath, content string) bool {
	_ = os.MkdirAll(path.Dir(filepath), 0777)
	if _, err := ioutil.ReadFile(filepath); err != nil {
		err := ioutil.WriteFile(filepath, []byte(content), 0777)
		if err == nil {
			fmt.Printf("文件：【%s】 生成成功！\n", filepath)
			return true
		}
	} else {
		fmt.Printf("文件：【%s】 已经存在！\n", filepath)
	}
	return false
}

// 获取适合linux和window的路径
func GetFilePath(elem ...string) string {

	for i, v := range elem {
		if strings.Index(v, "app\\") == 0 {
			v = strings.Replace(v, "app\\", "app\\", 1)
		}
		elem[i] = strings.ReplaceAll(v, "\\", "/")
	}
	return path.Join(elem...)
}

// 驼峰转下划线/-
func Camel2Case(ori string, want byte) string {
	newStr := make([]byte, 0)
	for i, r := range ori {
		if unicode.IsUpper(r) {
			if i != 0 {
				newStr = append(newStr, want)
			}
			newStr = append(newStr, byte(unicode.ToLower(r)))
		} else {
			newStr = append(newStr, byte(r))
		}
	}
	return strings.ToLower(string(newStr))
}

// Capitalize 字符首字母大写
func CapOrLow(str string, cap bool) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if cap && vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else if cap == false && vv[i] >= 65 && vv[i] <= 90 {
				vv[i] += 32
				upperStr += string(vv[i])
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
