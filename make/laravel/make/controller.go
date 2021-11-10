package make

import (
	"strings"
	"yii/make/laravel/template"
	"yii/make/laravel/util"
	util_tp "yii/make/tp/util"
)

type Controller struct {
	Name       string
	Namespace  string
	Root       string
	Table      string
	NameSpaces []string
	DbPrefix   string
}

func (m *Controller) Make() {
	filepath := util.GetFilePath(m.Root, m.Namespace, m.Name+"Controller.php")
	content := strings.ReplaceAll(template.Controller, "{$namespace}", m.Namespace)
	content = strings.ReplaceAll(content, "{$name}", m.Name)
	content = strings.ReplaceAll(content, "{$ucName}", util.CapOrLow(m.Name, false))
	util_tp.WriteFile(filepath, content)
	// 生成repository和route
	m.otherMake()
}

func (m *Controller) otherMake() {
	name_ := util.Camel2Case(m.Name, '_')
	last := m.NameSpaces[len(m.NameSpaces)-1]
	last = util.CapOrLow(last, false)
	routePath := util.GetFilePath(m.Root, "routes", name_+".php")
	if last != "" {
		routePath = util.GetFilePath(m.Root, "routes", last, name_+".php")
	}
	route := &Route{
		Namespace:  m.Namespace,
		PathName:   routePath,
		Prefix:     name_,
		Controller: m.Name,
	}
	route.Make()
	reNameSpace := "App\\Repository"
	reNamespaces := strings.Split(reNameSpace, "\\")
	re := &Repository{
		Name:       m.Name,
		Namespace:  reNameSpace,
		Root:       m.Root,
		NameSpaces: reNamespaces,
		DbPrefix:   m.DbPrefix,
	}
	re.Make()
}
