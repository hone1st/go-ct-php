package make

import (
	"honelst/util/mysql"
	"honelst/util/other"
	"log"
	"os"
)

type Curd struct {
	Name       string
	NameSpace  string
	NameSpaces []string
	Table      string
	Root       string
}

func (c Curd) Make() {
	//c.Root = "J:\\bitch_admin"
	// 创建模块
	// 读取数据库的数据
	tableName := other.Camel2Case(c.Table, '_')
	db, error := mysql.MysqlConnect(mysql.LoadConfigEnvTp(c.Root))
	if error != nil {
		log.Fatal(error.Error())
	}
	fieldsMap, fields, tableComment := mysql.TableFieldsMap(db, tableName)
	lModel := other.CapOrLow(c.Table, false)
	model := c.Table
	module := c.NameSpaces[1]
	controller := other.GetFilePath(c.Root, "application", module, "admin", model+".php")
	if _, err := os.Stat(controller); err != nil {
		AdminController{
			LModel: lModel,
			Model:  model,
			Module: module,
			Path:   controller,
		}.Make()
	}
	validate := other.GetFilePath(c.Root, "application", module, "validate", model+".php")
	if _, err := os.Stat(validate); err != nil {
		AdminValidate{
			Model:  model,
			Module: module,
			Mp:     fieldsMap,
			Fields: fields,
			Path:   validate,
		}.Make()
	}
	modelPath := other.GetFilePath(c.Root, "application", module, "model", model+".php")
	if _, err := os.Stat(modelPath); err != nil {
		AdminModel{
			Model:        model,
			Module:       module,
			Mp:           fieldsMap,
			Fields:       fields,
			Path:         modelPath,
			TableComment: tableComment,
		}.Make()
	}
	service := other.GetFilePath(c.Root, "application", module, "service", model+".php")
	if _, err := os.Stat(service); err != nil {
		AdminService{
			Model:  model,
			Module: module,
			Path:   service,
		}.Make()
	}
	view := other.GetFilePath(c.Root, "application", module, "view", tableName)
	AdminTemplate{
		Model:  model,
		Module: module,
		Mp:     fieldsMap,
		Fields: fields,
		Dir:    view,
	}.Make()

}
