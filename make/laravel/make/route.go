package make

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"yii/make/laravel/template"
	util_tp "yii/make/tp/util"
)

type Route struct {
	Namespace  string
	PathName   string
	Prefix     string
	Controller string

	ControllerFilePath string
}

func (m *Route) Make() {
	content := strings.ReplaceAll(template.Route, "{$prefix}", m.Prefix)
	content = strings.ReplaceAll(content, "{$controller}", m.Controller)
	content = strings.ReplaceAll(content, "{$namespace}", m.Namespace)
	util_tp.WriteFile(m.PathName, content)
	m.update()
}

func (m *Route) update() {
	if m.ControllerFilePath == "" {
		return
	}
	// 先匹配控制器
	routeRe, _ := regexp.Compile("public function [a-z][a-zA-Z0-9_]+\\(")
	controllerB, _ := ioutil.ReadFile(m.ControllerFilePath)
	routeResult := routeRe.FindAllString(string(controllerB), -1)
	existB, _ := ioutil.ReadFile(m.PathName)
	exist := string(existB)
	writeArr := make([]string, 0)
	for i := 0; i < len(routeResult); i++ {
		routeResult[i] = strings.ReplaceAll(routeResult[i], "public function ", "")
		routeResult[i] = strings.ReplaceAll(routeResult[i], "(", "")
		check := fmt.Sprintf("%sController::class, '%s'", m.Controller, routeResult[i])
		existRe, _ := regexp.Compile(check)
		if len(existRe.FindAllString(exist, -1)) > 0 {
			continue
		}
		addRoute := fmt.Sprintf("    Route::get('%s', [%s]);", routeResult[i], check)
		log.Println("即将更新路由新增：", addRoute)
		writeArr = append(writeArr, addRoute)
	}
	if len(writeArr) == 0 {
		return
	}
	endRegex, _ := regexp.Compile("[\r\n]+}\\);")
	endResult := endRegex.FindString(exist)
	if endResult == "" {
		log.Println("路由文件更新失败! 格式不对")
		return
	}
	content := strings.ReplaceAll(exist, endResult, "\r\n"+strings.Join(writeArr, "\r\n")+"\r\n});")
	_ = ioutil.WriteFile(m.PathName, []byte(content), 0777)
	log.Println("路由文件更新成功!")
}
