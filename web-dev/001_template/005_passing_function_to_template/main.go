package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var absPath = `/Users/debs/Personal/start_with_go/web-dev/001_template/005_passing_function_to_template`

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThreeLetters,
}

func firstThreeLetters(s string ) string{
	s = strings.TrimSpace(s)
	return s[:3]
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.New("xtemplate.gohtml").Funcs(fm).ParseFiles(absPath+"/xtemplate.gohtml"))
}
func main() {
	newFile, err := os.Create(absPath+"/index.html")
	if err!=nil{
		log.Fatal(err)
	}
	data := "Hello there How are you"
	err = tpl.Execute(newFile, data)
	if err!=nil{
		log.Fatal(err)
	}
}
