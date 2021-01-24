package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template
var absPath string = "/Users/debs/Personal/start_with_go/web-dev/001_template/002_text_template_parseGlob"

func init(){
	tpl = template.Must(template.ParseGlob(absPath+"/*.gohtml"))
}

func main() {
	newFile, err := os.Create(absPath+"/index.html")
	if err!=nil{
		log.Fatal(err)
	}
	err = tpl.Execute(newFile, nil)
	if err!=nil{
		log.Fatal(err)
	}
}
