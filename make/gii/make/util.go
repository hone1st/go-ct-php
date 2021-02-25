/**
 * @Time : 2020/10/9 16:54
 * @Author : liang
 * @File : util
 * @Software: GoLand
 */

package make

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

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

// 根据连接信息生成model数据
func GetPropertyByTable(tableName, root string) string {
	db, dbName := initDb(root)
	if db == nil {
		return ""
	}
	res, err := db.Query("select column_name,column_comment,data_type from "+
		"information_schema.columns where table_schema =? and table_name = ? ; ", dbName, tableName)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	defer res.Close()
	str := "/**\n"
	for res.Next() {
		var columnName, columnComment, dataType string
		_ = res.Scan(&columnName, &columnComment, &dataType)
		dataType = typeToPhp(dataType)
		str = str + fmt.Sprintf(" * @property %s $%s %s\n", dataType, columnName, columnComment)
	}
	if str == "/**\n" {
		return ""
	}
	return str + " */"
}

const Basic = "int,tinyint,bigint"

// 数据库转化为php的类型
func typeToPhp(t string) string {
	phpType := "string"
	if strings.Contains(Basic, strings.ToLower(t)) {
		phpType = "int"
	}
	return phpType
}

// 获取适合linux和window的路径
func GetFilePath(elem ...string) string {
	for i, v := range elem {
		elem[i] = strings.ReplaceAll(v, "\\", "/")
	}
	return path.Join(elem...)
}

func ReplaceMoreStr(ori, new string, old ...string) string {
	for _, v := range old {
		ori = strings.ReplaceAll(ori, v, new)
	}
	return ori
}

// 读取配置获取db
func initDb(root string) (*sql.DB, string) {
	var (
		host, dbName, username, password string
	)
	configPath := GetFilePath(root, "common/config/main-local.php")
	config, _ := ioutil.ReadFile(configPath)
	// 匹配  'db' =>... ],
	re := regexp.MustCompile("'db'[^~]*?\\]")
	phpConfig := re.FindString(string(config))
	// 获取dsn username password host dbname
	dsnRe := regexp.MustCompile("'mysql[^~]*?'")
	usernameRe := regexp.MustCompile("'username[^~]*?,")
	passwordRe := regexp.MustCompile("'password[^~]*?,")

	dsn := dsnRe.FindString(phpConfig)
	usernameStr := usernameRe.FindString(phpConfig)
	passwordStr := passwordRe.FindString(phpConfig)

	dsns := strings.Split(dsn, ":")
	if len(dsns) != 3 {
		return nil, ""
	}
	host = strings.Split(dsns[1], "=")[1]
	dbName = strings.Split(dsns[2], "=")[1]

	dbName = ReplaceMoreStr(dbName, "", "'", " ")
	username = ReplaceMoreStr(usernameStr, "", "'", "username", " ", "=>", ",")
	password = ReplaceMoreStr(passwordStr, "", "'", "password", " ", "=>", ",")

	Db, DbErr := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", username, password, host, dbName))
	if DbErr != nil {
		log.Println("数据库连接失败！")
		return nil, ""
	}
	return Db, dbName
}
