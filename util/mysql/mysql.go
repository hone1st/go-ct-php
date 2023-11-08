package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gopkg.in/ini.v1"
	"honelst/util/other"
	"log"
	"os"
	"path"
	"strconv"
)

type MysqlConfig struct {
	host     string
	username string
	password string
	database string
	prefix   string
	port     int
}

type DbResult struct {
	Db     *sql.DB
	Config *MysqlConfig
}

// LoadConfigEnvLaravel 加载根目录的env
func LoadConfigEnvLaravel(root, dbPrefix string) *MysqlConfig {
	err := godotenv.Load(path.Join(root, ".env"))
	if err != nil {
		log.Fatal(err)
	}
	host := os.Getenv(dbPrefix + "DB_HOST")
	portStr := os.Getenv(dbPrefix + "DB_PORT")
	port, _ := strconv.Atoi(portStr)
	database := os.Getenv(dbPrefix + "DB_DATABASE")
	username := os.Getenv(dbPrefix + "DB_USERNAME")
	password := os.Getenv(dbPrefix + "DB_PASSWORD")
	prefix := os.Getenv(dbPrefix + "DB_PREFIX")

	// 二次处理加密的数据
	if deName := other.DecryptData(username); deName != "" {
		username = deName
	}
	if dePassword := other.DecryptData(password); dePassword != "" {
		password = dePassword
	}

	return &MysqlConfig{
		host:     host,
		username: username,
		password: password,
		database: database,
		port:     port,
		prefix:   prefix,
	}
}

// LoadConfigEnvTp 加载根目录的env
func LoadConfigEnvTp(root string) *MysqlConfig {
	cfg, err := ini.Load(path.Join(root, ".env"))
	if err != nil {
		log.Println(".env文件不存在", err)
		os.Exit(0)
	}
	host := cfg.Section("").Key("DB_HOST").String()
	port, _ := cfg.Section("").Key("DB_PORT").Int()
	database := cfg.Section("").Key("DB_DATABASE").String()
	username := cfg.Section("").Key("DB_USERNAME").String()
	password := cfg.Section("").Key("DB_PASSWORD").String()

	prefix := cfg.Section("").Key("DB_PREFIX").String()
	return &MysqlConfig{
		host:     host,
		username: username,
		password: password,
		database: database,
		port:     port,
		prefix:   prefix,
	}
}

// MysqlConnect 获取数据库的链接
func MysqlConnect(config *MysqlConfig) (*DbResult, error) {
	db, DbErr := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		config.username, config.password, config.host, config.port, config.database))
	if DbErr != nil {
		log.Println("数据库连接失败！")
		return nil, DbErr
	}
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	return &DbResult{
		Db:     db,
		Config: config,
	}, nil
}
