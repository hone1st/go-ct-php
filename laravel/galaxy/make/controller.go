package make

import (
	"honelst/laravel/galaxy/template"
	"honelst/util/other"
	"strings"
)

type Controller struct {
	Name       string
	Namespace  string
	Root       string
	Table      string
	NameSpaces []string
	DbPrefix   string

	FilePath string
}

func (m *Controller) Make() {
	filepath := other.GetFilePath(m.Root, m.Namespace, m.Name+"Controller.php")
	m.FilePath = filepath
	content := strings.ReplaceAll(template.Controller, "{$namespace}", m.Namespace)
	content = strings.ReplaceAll(content, "{$name}", m.Name)
	content = strings.ReplaceAll(content, "{$ucName}", other.CapOrLow(m.Name, false))
	other.WriteFile(filepath, content)
	// 生成repository和route
	m.otherMake()
}

func (m *Controller) otherMake() {
	name_ := other.Camel2Case(m.Name, '_')
	last := m.NameSpaces[len(m.NameSpaces)-1]
	last = other.CapOrLow(last, false)
	routePath := other.GetFilePath(m.Root, "routes", name_+".php")
	if last != "" {
		routePath = other.GetFilePath(m.Root, "routes", last, name_+".php")
	}
	route := &Route{
		Namespace:          m.Namespace,
		PathName:           routePath,
		Prefix:             other.Camel2Case(m.Name, '-'),
		Controller:         m.Name,
		ControllerFilePath: m.FilePath,
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

	va := &Validation{
		Name:       m.Name,
		Namespace:  m.Namespace,
		Root:       m.Root,
		NameSpaces: m.NameSpaces,
		DbPrefix:   m.DbPrefix,
	}
	va.Make()
}
