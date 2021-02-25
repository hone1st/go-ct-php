/**
 * @Time : 2020/10/9 16:28
 * @Author : liang
 * @File : make
 * @Software: GoLand
 */

package make

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"yii/make/gii/template"
)

type Model struct {
	Name      string
	NameSpace string
	Root      string
}

func (m *Model) Make() {
	tableName := Camel2Case(m.Name, '_')
	property := GetPropertyByTable(tableName, m.Root)
	filepath := GetFilePath(m.Root, m.NameSpace, m.Name+".php")
	content := fmt.Sprintf(template.Model, m.NameSpace, property, m.Name, tableName)
	if !WriteFile(filepath, content) && property != "" {
		re := regexp.MustCompile(fmt.Sprintf("/\\*\\*[^`]*?class %s", m.Name))
		ori, _ := ioutil.ReadFile(filepath)
		dest := strings.ReplaceAll(string(ori), re.FindString(string(ori)), fmt.Sprintf("%s\nclass %s", property, m.Name))
		if dest != string(ori) {
			_ = ioutil.WriteFile(filepath, []byte(dest), 0777)
		}
	}
}

type Controller struct {
	Name       string
	NameSpace  string
	NameSpaces []string
	Root       string
	Reset      bool
}

func (c *Controller) Make() {
	diController := `use common\controllers\DIController;`
	parentController := `DiController`
	switch c.NameSpaces[0] {
	case "console":
		diController = `namespace console\controllers;`
		parentController = `Controller`
		break
	}
	filepath := GetFilePath(c.Root, c.NameSpace, c.Name+"Controller.php")
	content := fmt.Sprintf(template.Controller, c.NameSpace, diController, c.Name, parentController)
	if c.Reset {
		content = fmt.Sprintf(template.Reset, c.NameSpace, diController, c.Name, parentController)
	}
	WriteFile(filepath, content)
}

type Service struct {
	Name      string
	NameSpace string
	Root      string
}

func (s *Service) Make() {
	serviceContent := fmt.Sprintf(template.Service, s.NameSpace, s.Name)
	serviceImplContent := fmt.Sprintf(template.ServiceImpl, s.NameSpace, s.NameSpace+"\\"+s.Name, s.Name, s.Name)
	serviceFilepath := GetFilePath(s.Root, s.NameSpace, s.Name+"Service.php")
	serviceImplFilepath := GetFilePath(s.Root, s.NameSpace, "impl", s.Name+"ServiceImpl.php")
	if WriteFile(serviceFilepath, serviceContent) && WriteFile(serviceImplFilepath, serviceImplContent) {
		diPath := GetFilePath(s.Root, "common", "config", "di-container.php")
		diContent, err := ioutil.ReadFile(diPath)
		if err != nil {
			log.Println("di-container文件不存在！")
		}
		appendContent := fmt.Sprintf("Yii::$container->set(\\%s\\%sService::class, \\%s\\impl\\%sServiceImpl::class);",
			s.NameSpace, s.Name, s.NameSpace, s.Name)
		if !strings.Contains(string(diContent), appendContent) {
			_ = ioutil.WriteFile(diPath, []byte(string(diContent)+"\n"+appendContent), 0777)
		}
	}
}

type Form struct {
	Name      string
	NameSpace string
	Root      string
}

func (f *Form) Make() {
	content := fmt.Sprintf(template.Form, f.NameSpace, f.Name)
	filepath := GetFilePath(f.Root, f.NameSpace, f.Name+"Form.php")
	WriteFile(filepath, content)
}

type Module struct {
	Name       string
	NameSpace  string
	Root       string
	NameSpaces []string
}

// 生成module
func (m *Module) Make() {
	// 不生成最顶级的module 所以命名空间必须存在两个modules
	index := 0
	diNamespaces := make([]string, 0)
	for i, v := range m.NameSpaces {
		if v == "modules" {
			index++
		}
		if index >= 2 {
			diNamespaces = append(diNamespaces, m.NameSpaces[i:]...)
			break
		}
	}
	if index <= 1 {
		log.Println("不生成根模块")
		os.Exit(1)
	}
	parent := strings.Join(append(m.NameSpaces[:len(m.NameSpaces)-1], "Module"), "\\")
	filepath := GetFilePath(m.Root, m.NameSpace, m.Name, "Module.php")
	content := fmt.Sprintf(template.Module, m.NameSpace+"\\"+m.Name, "\\"+parent)
	if WriteFile(filepath, content) {
		parentC, _ := ioutil.ReadFile(GetFilePath(m.Root, parent+".php"))
		dest := fmt.Sprintf("%s\\%s\\Module::class", strings.Join(diNamespaces, "\\"), m.Name)
		if !strings.Contains(string(parentC), dest) {
			re := regexp.MustCompile("this->modules[^~]*\\];")
			ori := re.FindString(string(parentC))
			newC :=
				strings.Replace(
					string(parentC),
					ori,
					strings.Replace(ori, "];", fmt.Sprintf("    '%s' => %s,\n        ];", Camel2Case(m.Name, '-'), dest), 1),
					1)
			_ = ioutil.WriteFile(GetFilePath(m.Root, parent+".php"), []byte(newC), 0777)
		}
	}
}
