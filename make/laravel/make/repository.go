package make

import (
	"strings"
	"yii/make/laravel/template"
	"yii/make/laravel/util"
	util_tp "yii/make/tp/util"
)

type Repository struct {
	Name       string
	Namespace  string
	Root       string
	Table      string
	NameSpaces []string
	DbPrefix   string
}

func (m *Repository) Make() {
	filepath := util.GetFilePath(m.Root, m.Namespace, m.Name+"Repository.php")
	content := strings.ReplaceAll(template.Repository, "{$namespace}", m.Namespace)
	content = strings.ReplaceAll(content, "{$name}", m.Name)
	util_tp.WriteFile(filepath, content)
}
