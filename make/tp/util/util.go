package util_tp

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"unicode"
)

type Field struct {
	ColumnName    string `json:"column_name"`
	ColumnComment string `json:"column_comment"`
	DataType      string `json:"data_type"`
}

type DbResult struct {
	Db       *sql.DB
	Database string
	Prefix   string
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

func TableFieldsMap(root, tableName string, namespaces []string) (map[string]*Field, []*Field, string) {
	db := initDb(root, namespaces)
	mp := make(map[string]*Field, 0)
	fields := make([]*Field, 0)
	if db == nil {
		return mp, fields, ""
	}
	res, err := db.Db.Query("select column_name,column_comment,data_type from "+
		"information_schema.columns where table_schema =? and table_name = ? ; ", db.Database, db.Prefix+tableName)
	if err != nil {
		log.Println(err)
	}
	tableRes, err := db.Db.Query("select table_comment from "+
		"information_schema.TABLES where table_schema =? and table_name = ? ; ", db.Database, db.Prefix+tableName)
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
	for res.Next() {
		_ = res.Scan(&columnName, &columnComment, &dataType)
		f := &Field{
			ColumnName:    columnName,
			ColumnComment: columnComment,
			DataType:      dataType,
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
	result := strings.Split(reContent, "=>")
	dest := result[1]
	dest = strings.ReplaceAll(dest, "'", "")
	dest = strings.ReplaceAll(dest, ",", "")
	return strings.TrimSpace(dest)
}

// 读取配置获取db
func initDb(root string, namespaces []string) *DbResult {
	var (
		host, dbName, username, password, prefix string
	)
	// 先找模块下的配置文件不存在的话就去找主配置的
	module := fmt.Sprintf("./application/%s/config/database.php", namespaces[1])

	var config []byte
	var errFile error
	if config, errFile = ioutil.ReadFile(path.Join(root, module)); errFile == nil {
		log.Printf("当前读取的数据库配置文件的路径是：【%s】\n", path.Join(root, module))
	} else if config, errFile = ioutil.ReadFile(path.Join(root, "./application/config/database.php")); errFile == nil {
		log.Printf("当前读取的数据库配置文件的路径是：【%s】\n", path.Join(root, "./application/config/database.php"))
	} else {
		config, _ = ioutil.ReadFile(path.Join(root, "./config/database.php"))
		log.Printf("当前读取的数据库配置文件的路径是：【%s】\n", path.Join(root, "./config/database.php"))
	}
	// 获取hostname/database/username/password
	hostnameRe := regexp.MustCompile("'hostname[^~]*?,")
	databaseRe := regexp.MustCompile("'database[^~]*?,")
	usernameRe := regexp.MustCompile("'username[^~]*?,")
	passwordRe := regexp.MustCompile("'password[^~]*?,")
	prefixRe := regexp.MustCompile("'prefix'[^~]*?,")

	prefix = getValueByRe(prefixRe.FindString(string(config)))
	host = getValueByRe(hostnameRe.FindString(string(config)))
	dbName = getValueByRe(databaseRe.FindString(string(config)))
	username = getValueByRe(usernameRe.FindString(string(config)))
	password = getValueByRe(passwordRe.FindString(string(config)))

	db, DbErr := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", username, password, host, dbName))
	if DbErr != nil {
		log.Println("数据库连接失败！")
		return nil
	}
	err := db.Ping()
	if err != nil {
		log.Println("数据库连接失败！正在尝试从根目录的.env获取配置")
		// 从新获取host/username/password
		fmt.Println(path.Join(root, ".env"))
		cfg, err := ini.Load(path.Join(root, ".env"))
		if err != nil {
			log.Println(".env文件不存在", err)
			os.Exit(0)
		}
		host = cfg.Section("DEV").Key("database_hostname").String()
		username = cfg.Section("DEV").Key("database_username").String()
		password = cfg.Section("DEV").Key("database_password").String()
		db, _ = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", username, password, host, dbName))
	}
	return &DbResult{
		Db:       db,
		Database: dbName,
		Prefix:   prefix,
	}
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
			v = strings.Replace(v, "app\\", "application\\", 1)
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
