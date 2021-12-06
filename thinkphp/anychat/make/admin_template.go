package make

import (
	"crypto/md5"
	"fmt"
	"honelst/thinkphp/anychat/template"
	"honelst/util/mysql"
	"honelst/util/other"
	"io"
	"strings"
	"time"
)

type AdminTemplate struct {
	Model  string
	Module string
	Mp     map[string]*mysql.Field
	Fields []*mysql.Field
	Dir    string
}

func (t AdminTemplate) Make() {
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
func (t AdminTemplate) update(content []string) {
	fileContent := template.Update
	path := other.GetFilePath(t.Dir, "update.html")
	fileContent = strings.ReplaceAll(fileContent, "{$update}", strings.Join(content, "\r\n"))
	other.WriteFile(path, fileContent)
}

func (t *AdminTemplate) noPass() {
	fileContent := template.NoPass
	path := other.GetFilePath(t.Dir, "no_pass.html")
	other.WriteFile(path, fileContent)
}

func (t AdminTemplate) create(content []string) {
	fileContent := template.Create
	path := other.GetFilePath(t.Dir, "create.html")
	fileContent = strings.ReplaceAll(fileContent, "{$create}", strings.Join(content, "\r\n"))
	other.WriteFile(path, fileContent)
}

func (t AdminTemplate) index(content []string) {
	fileContent := template.AdminIndex
	path := other.GetFilePath(t.Dir, "index.html")
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
	other.WriteFile(path, fileContent)
}
