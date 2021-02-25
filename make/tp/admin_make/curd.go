package admin_make

import (
	"os"
	util_tp "yii/make/tp/util"
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
	tableName := util_tp.Camel2Case(c.Table, '_')
	fieldsMap, fields, tableComment := util_tp.TableFieldsMap(c.Root, tableName, c.NameSpaces)
	lModel := util_tp.CapOrLow(c.Table, false)
	model := c.Table
	module := c.NameSpaces[1]
	controller := util_tp.GetFilePath(c.Root, "application", module, "admin", model+".php")
	if _, err := os.Stat(controller); err != nil {
		Controller{
			LModel: lModel,
			Model:  model,
			Module: module,
			Path:   controller,
		}.Make()
	}
	validate := util_tp.GetFilePath(c.Root, "application", module, "validate", model+".php")
	if _, err := os.Stat(validate); err != nil {
		Validate{
			Model:  model,
			Module: module,
			Mp:     fieldsMap,
			Fields: fields,
			Path:   validate,
		}.Make()
	}
	modelPath := util_tp.GetFilePath(c.Root, "application", module, "model", model+".php")
	if _, err := os.Stat(modelPath); err != nil {
		Model{
			Model:        model,
			Module:       module,
			Mp:           fieldsMap,
			Fields:       fields,
			Path:         modelPath,
			TableComment: tableComment,
		}.Make()
	}
	service := util_tp.GetFilePath(c.Root, "application", module, "service", model+".php")
	if _, err := os.Stat(service); err != nil {
		Service{
			Model:  model,
			Module: module,
			Path:   service,
		}.Make()
	}
	view := util_tp.GetFilePath(c.Root, "application", module, "view", tableName)
	Template{
		Model:  model,
		Module: module,
		Mp:     fieldsMap,
		Fields: fields,
		Dir:    view,
	}.Make()

}
