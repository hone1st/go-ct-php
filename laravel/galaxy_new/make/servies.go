package make

import (
	"honelst/laravel/galaxy_new/template"
	"honelst/util/other"
	"strings"
)

type Services struct {
	Name       string
	Namespace  string
	Root       string
	Table      string
	NameSpaces []string
	DbPrefix   string
}

func (m *Services) Make() {
	filepath := other.GetFilePath(m.Root, m.Namespace, m.Name+"Services.php")
	content := strings.ReplaceAll(template.Services, "{$namespace}", other.CapOrLow(m.Namespace, true))
	content = strings.ReplaceAll(content, "{$namespace_parent}", other.CapOrLow(strings.Replace(m.Namespace, "\\Services", "", 1), true))
	content = strings.ReplaceAll(content, "{$name}", m.Name)
	content = strings.ReplaceAll(content, "{$ucName}", other.CapOrLow(m.Name, false))
	tableName := other.Camel2Case(m.Name, '_')
	content = strings.ReplaceAll(content, "{$name_id}", tableName+"_id")
	other.WriteFile(filepath, content)
}
