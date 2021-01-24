package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

var absPath string = `/Users/debs/Personal/start_with_go/web-dev/001_template/001_text_template`
func init(){

	tpl = template.Must(template.ParseFiles(absPath+"/my_first_go_template.gohtml"))

}
func main() {
	newFile, err:= os.Create(absPath+"/index.html")
	if err!=nil{
		log.Fatal(err)
	}
	tpl.Execute(newFile, nil)
}
