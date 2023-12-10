package mysql

import (
	"log"
	"strings"
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
	case "varchar", "char":
		return "string"
	case "int", "tinyint":
		return "int"
	case "decimal":
		return "float"
	case "json":
		return "array|null"
	case "datetime", "timestamp":
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
		"information_schema.columns where table_schema =? and table_name = ? order by ordinal_position asc; ", db.Config.database, db.Config.prefix+tableName)
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
