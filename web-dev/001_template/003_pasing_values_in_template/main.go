package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template
var absPath string = `/Users/debs/Personal/start_with_go/web-dev/001_template/003_pasing_values_in_template`

func init(){
	tpl = template.Must(template.ParseFiles(absPath+"/passing_value.gohtml"))
}
func main() {
	newFile, err:= os.Create(absPath+"/index.html")
	if err!=nil{
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(newFile, "passing_value.gohtml", "SAMRAT")
	if err!=nil{
		log.Fatal(err)
	}


}
