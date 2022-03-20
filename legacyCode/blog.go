package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type contentHtml struct {
	Title   string
	Content []string
}

func handle(value string) string {
	return fmt.Sprintf(value + "yes")
}

func HelloTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.New("hello")                  //定义新模板
	t.Funcs(template.FuncMap{"handle": handle}) //隐射模板内管道函数
	t, err := t.Parse(`
	<head>
	<title>{{.Title}}</title>
	</head>
	<body>
	{{if .Content}}
	{{range .Content}}
	<h1>{{. | handle}}</h1>
	{{end}}
	{{else}}
	<p>No Content</p>
	{{end}}
	</body>
	</html>`)

	if err != nil {
		panic(fmt.Sprintf("template fail:%s", err.Error()))
	}
	text := contentHtml{
		Title:   "Hello Golang Blog",
		Content: []string{"Golang", "Java", "Python", "Rust", "C++"},
	}
	err = t.Execute(w, text) //装载模板
	if err != nil {
		panic(err)
	}
}
func main() {
	http.HandleFunc("/", HelloTemplate)
	log.Fatal(http.ListenAndServe(":80", nil))
}
