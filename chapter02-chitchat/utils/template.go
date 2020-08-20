package utils

import (
	"fmt"
	"html/template"
	"net/http"
	
	"k8s.io/klog"
)

// parse HTML templates
// pass in a list of file names, and get a template
func ParseTemplateFiles(filenames ...string) *template.Template {
	var files []string
	t := template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("chapter02-chitchat/static/templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return t
}

func GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("chapter02-chitchat/static/templates/%s.html", file))
	}
	
	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(writer, "layout", data)
	if err != nil {
		klog.Errorf("execute template error: %v", err)
	}
}
