package handlers

import (
	"html/template"
	"net/http"
)

type TestStruct struct {
	First  string
	Second string
}

func PrimHandler(r http.ResponseWriter, w *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	dummyData := []TestStruct{
		{First: "first struct", Second: "first struct"},
		{First: "seconds struct", Second: "sec struct"},
	}

	tmpl.Execute(r, dummyData)

}
