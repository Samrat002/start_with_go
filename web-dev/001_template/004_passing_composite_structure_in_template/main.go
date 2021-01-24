package main

import (
	"html/template"
	"log"
	"os"
)

var absPath string = `/Users/debs/Personal/start_with_go/web-dev/001_template/004_passing_composite_structure_in_template`
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles(absPath + `/my_html_template.gohtml`))

}



func main() {
	newFile, err := os.Create(absPath + "/index.html")
	if err != nil {
		log.Fatal(err)
	}
	data := []struct {
		Name    string
		Address string
		Zip     int64
		City    string
		Region  string
	}{
		{
			Name: "hotel1",
			Address: "streat address",
			Zip: 456789,
			City: "test city",
			Region: "NE",
		},
		{
			Name: "hotel2",
			Address: "streat address 1",
			Zip: 4567891,
			City: "test city differnt",
			Region: "NE",
		},

	}
	err = tpl.Execute(newFile, data)
	if err != nil {
		log.Fatal(err)
	}
}
