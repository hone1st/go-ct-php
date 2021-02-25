package admin_make

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"time"
	"yii/make/tp/admin_template"
	util_tp "yii/make/tp/util"
)

type Template struct {
	Model  string
	Module string
	Mp     map[string]*util_tp.Field
	Fields []*util_tp.Field
	Dir    string
}

func (t Template) Make() {
	update := make([]string, 0)
	create := make([]string, 0)
	index := make([]string, 0)
	if len(t.Mp) > 0 {
		for k, v := range t.Mp {
			create = append(create, fmt.Sprintf(`	<div class="layui-form-item">
            <label class="layui-form-label">%s</label>
            <div class="layui-input-block">
                <input type="text" class="layui-input field-name" name="%s" lay-verify="required"
                       autocomplete="off" placeholder="%s">
            </div>
     </div>`, v.ColumnComment, k, v.ColumnComment))
			update = append(update, fmt.Sprintf(`	<div class="layui-form-item">
            <label class="layui-form-label">%s</label>
            <div class="layui-input-block">
                <input type="text" class="layui-input field-name" value="{$item['%s']}" name="%s" lay-verify="required"
                       autocomplete="off" placeholder="%s">
            </div>
     </div>`, v.ColumnComment, k, k, v.ColumnComment))
			index = append(index, fmt.Sprintf(`				{field: '%s', title: '%s', align: 'center'}`, k, v.ColumnComment))
		}
	}
	t.update(update)
	t.create(create)
	t.index(index)
}
func (t Template) update(content []string) {
	fileContent := admin_template.Update
	path := util_tp.GetFilePath(t.Dir, "update.html")
	fileContent = strings.ReplaceAll(fileContent, "{$update}", strings.Join(content, "\r\n"))
	util_tp.WriteFile(path, fileContent)
}

func (t Template) create(content []string) {
	fileContent := admin_template.Create
	path := util_tp.GetFilePath(t.Dir, "create.html")
	fileContent = strings.ReplaceAll(fileContent, "{$create}", strings.Join(content, "\r\n"))
	util_tp.WriteFile(path, fileContent)
}

func (t Template) index(content []string) {
	fileContent := admin_template.Index
	path := util_tp.GetFilePath(t.Dir, "index.html")
	if len(content) > 0 {
		fileContent = strings.ReplaceAll(fileContent, "{$cols}", strings.Join(content, ",\r\n")+",")
	} else {
		fileContent = strings.ReplaceAll(fileContent, "{$cols}", "")
	}
	w := md5.New()
	_, _ = io.WriteString(w, time.Now().Format(time.RFC3339))
	//将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	fileContent = strings.ReplaceAll(fileContent, "{$uuid}", md5str2)
	util_tp.WriteFile(path, fileContent)
}
