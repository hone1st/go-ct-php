package make

import (
	"strings"
	"yii/make/laravel/template"
	util_tp "yii/make/tp/util"
)

type Route struct {
	Namespace  string
	PathName   string
	Prefix     string
	Controller string
}

func (m *Route) Make() {
	content := strings.ReplaceAll(template.Route, "{$prefix}", m.Prefix)
	content = strings.ReplaceAll(content, "{$controller}", m.Controller)
	content = strings.ReplaceAll(content, "{$namespace}", m.Namespace)
	util_tp.WriteFile(m.PathName, content)
}
