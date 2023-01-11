package make

import (
	"honelst/laravel/galaxy_new/template"
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
	tableName := other.Camel2Case(m.Name, '_')
	content := strings.ReplaceAll(template.Controller, "{$namespace}", m.Namespace)
	content = strings.ReplaceAll(content, "{$namespace_parent}", strings.Replace(m.Namespace, "\\Controllers", "", 1))
	content = strings.ReplaceAll(content, "{$name}", m.Name)
	content = strings.ReplaceAll(content, "{$name_id}", tableName+"_id")
	content = strings.ReplaceAll(content, "{$ucName}", other.CapOrLow(m.Name, false))
	other.WriteFile(filepath, content)
	// 生成repository和route
	m.otherMake()
}

func (m *Controller) otherMake() {
	name_ := other.Camel2Case(m.Name, '_')
	parentNameSpaces := m.NameSpaces[:len(m.NameSpaces)-1]
	parentNameSpace := strings.Join(parentNameSpaces, "\\")
	last := parentNameSpaces[len(parentNameSpaces)-1]
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

	re := &Repository{
		Name:       m.Name,
		Namespace:  parentNameSpace + "\\Repository",
		Root:       m.Root,
		NameSpaces: strings.Split(parentNameSpace+"\\Repository", "\\"),
		DbPrefix:   m.DbPrefix,
	}
	re.Make()

	va := &Validation{
		Name:       m.Name,
		Namespace:  parentNameSpace + "\\Validations",
		Root:       m.Root,
		NameSpaces: strings.Split(parentNameSpace+"\\Validations", "\\"),
		DbPrefix:   m.DbPrefix,
	}
	va.Make()

	se := &Services{
		Name:       m.Name,
		Namespace:  parentNameSpace + "\\Services",
		Root:       m.Root,
		NameSpaces: strings.Split(parentNameSpace+"\\Services", "\\"),
		DbPrefix:   m.DbPrefix,
	}
	se.Make()
}
