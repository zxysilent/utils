package utils

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// LoadTmpl 加载指定目录解析为模板 仅限 .html 文件
// eg: views/index.html ,views/subtmpl/subtmpl.html
// tmpl:=LoadTmpl("./views",funcs)
// tmpl.ExecuteTemplate(yourWriter, "index.html", yourData)、tmpl.ExecuteTemplate(yourWriter, "subtmpl/subtmpl.html", yourData)
func LoadTmpl(root string, funcs template.FuncMap) *template.Template {
	tmpl := template.New("LoadTmpl")
	if funcs != nil {
		tmpl.Funcs(funcs)
	}
	rln := len(root)
	filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		pln := len(path)
		// 是文件 并且是 .html 结尾
		if !fi.IsDir() && pln > 4 && path[pln-5:] == ".html" {
			buf, err := ioutil.ReadFile(path)
			if err != nil {
				panic("read file :" + path + ". error :" + err.Error())
			}
			name := strings.ReplaceAll(path, "\\", "/")
			_, err = tmpl.New(name[rln-1:]).Parse(string(buf))
			if err != nil {
				panic("parse file :" + path + ". error :" + err.Error())
			}
		}
		return nil
	})
	return tmpl
}
