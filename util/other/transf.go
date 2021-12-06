package other

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"unicode"
)

// 获取适合linux和window的路径

func GetFilePath(elem ...string) string {

	for i, v := range elem {
		if strings.Index(v, "app\\") == 0 {
			v = strings.Replace(v, "app\\", "app\\", 1)
		}
		elem[i] = strings.ReplaceAll(v, "\\", "/")
	}
	return path.Join(elem...)
}

// 驼峰转下划线/-

func Camel2Case(ori string, want byte) string {
	newStr := make([]byte, 0)
	for i, r := range ori {
		if unicode.IsUpper(r) {
			if i != 0 {
				newStr = append(newStr, want)
			}
			newStr = append(newStr, byte(unicode.ToLower(r)))
		} else {
			newStr = append(newStr, byte(r))
		}
	}
	return strings.ToLower(string(newStr))
}

// Capitalize 字符首字母大写

func CapOrLow(str string, cap bool) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if cap && vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else if cap == false && vv[i] >= 65 && vv[i] <= 90 {
				vv[i] += 32
				upperStr += string(vv[i])
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// 生成文件

func WriteFile(filepath, content string) bool {
	_ = os.MkdirAll(path.Dir(filepath), 0777)
	if _, err := ioutil.ReadFile(filepath); err != nil {
		err := ioutil.WriteFile(filepath, []byte(content), 0777)
		if err == nil {
			fmt.Printf("文件：【%s】 生成成功！\n", filepath)
			return true
		}
	} else {
		fmt.Printf("文件：【%s】 已经存在！\n", filepath)
	}
	return false
}
