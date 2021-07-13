package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func staticFiles(w http.ResponseWriter, r *http.Request) {
	tempz, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
		tempz.Execute(w, nil)

	}
}
