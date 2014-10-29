package httpmain

import (
	//"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/add/search", addSearchHandler)
	http.HandleFunc("/", homeHandler)

}

type Page struct {
	Title string //Title of webpage
	Body  string //Body in byte form.
	URL   string //URL of the request

}
