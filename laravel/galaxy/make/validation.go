package make

import (
	"honelst/laravel/galaxy/template"
	"honelst/util/other"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// 构建验证层的代码 简单构建validation

type Validation struct {
	Name       string
	Namespace  string
	Root       string
	Table      string
	NameSpaces []string
	DbPrefix   string
}

func (v *Validation) Make() {
	// 固定路径
	// 检测文件是否存在
	controllerFilePath := other.GetFilePath(v.Root, v.Namespace, v.Name+"Controller.php")
	controllerData, err := ioutil.ReadFile(controllerFilePath)
	if err != nil {
		log.Fatalf("控制器文件不存在：【%s】 无法生成验证层文件", controllerFilePath)
	}
	controllerPath := other.GetFilePath(v.Root, v.Namespace)
	split := strings.Split(controllerPath, "App/Http/Controllers")

	if len(split) == 2 {
		validationPath := other.GetFilePath(v.Root, "app", "Http", "Validations", split[1], v.Name+"Validation.php")
		// 检测是否存在
		file, err := ioutil.ReadFile(validationPath)
		apiRegex := regexp.MustCompile("public function [a-z][a-zA-Z0-9_]+\\(")
		apiResult := apiRegex.FindAllString(string(controllerData), -1)
		if err != nil {
			// 创建新文件
			log.Printf("验证层文件已经不存在！即将创建文件：【%s】\r\n", validationPath)
			tem := template.Validation
			tem = strings.ReplaceAll(tem, "{$namespace}", strings.ReplaceAll(v.Namespace, "Controllers", "Validations"))
			tem = strings.ReplaceAll(tem, "{$name}", v.Name)
			v.makeFile(make([]string, 0), apiResult, tem, validationPath)
		} else {
			log.Printf("验证层文件已经存在！即将更新文件：【%s】\n", validationPath)
			// 存在就不创建新文件
			// 读取文件的接口文件
			validationResult := apiRegex.FindAllString(string(file), -1)
			v.makeFile(validationResult, apiResult, string(file), validationPath)
		}
		// 检测是否存在
	}

}

func (v *Validation) makeFile(validationResult, apiResult []string, ori, validationPath string) {
	mp := make(map[string]int, 0)
	for i := 0; i < len(apiResult); i++ {
		mp[apiResult[i]] = 1
		for j := 0; j < len(validationResult); j++ {
			if validationResult[j] == apiResult[i] {
				delete(mp, apiResult[i])
				break
			}
		}
	}
	if len(mp) > 0 {
		var template = `    /**
     * @see {$controller}::{$method}
     * @return \string[][]
     */
    {$function})
    {
        return [
            'rules' => [

            ],
            'messages' => [

            ],
        ];
    }

`
		write := make([]string, 0)
		for key := range mp {
			log.Printf("即将增加验证层的方法：%s \r\n", key)
			method := strings.ReplaceAll(key, "public function ", "")
			method = strings.ReplaceAll(method, "(", "")
			temp := template
			temp = strings.ReplaceAll(temp, "{$controller}", "\\"+v.Namespace+"\\"+v.Name+"Controller")
			temp = strings.ReplaceAll(temp, "{$method}", method)
			temp = strings.ReplaceAll(temp, "{$function}", key)
			write = append(write, temp)
		}
		//匹配结束位置  }//end
		endRegex := regexp.MustCompile("\\}//end")
		end := endRegex.FindString(ori)
		if end == "" {
			log.Fatalf("找不到结束表示位：【}//end】，无法更新!")
		}
		write = append(write, end)
		ori = strings.ReplaceAll(ori, end, strings.Join(write, "\r\n"))
		ioutil.WriteFile(validationPath, []byte(ori), 0777)
	}
}
