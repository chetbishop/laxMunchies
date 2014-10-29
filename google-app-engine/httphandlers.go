package httpmain

import (
	"html/template"
	"net/http"
)

func addSearchHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	p.Title = "Add Anime"
	if r.Method == "GET" {
		t, _ := template.ParseFiles("web/addAnimeSearch.html", "web/header.html", "web/footer.html")
		t.ExecuteTemplate(w, "header", p)
		t.ExecuteTemplate(w, "addAnimeSearch", p)
		t.ExecuteTemplate(w, "footer", p)
	} else if r.Method == "POST" {
		r.ParseForm()
		//aname := r.FormValue("animename")
		//p.Anime = anidbapi.AnimeSearchWrapper(runningConfig, aname)
		t, _ := template.ParseFiles("web/addAnimeResults.html", "web/header.html", "web/footer.html")
		t.ExecuteTemplate(w, "header", p)
		t.ExecuteTemplate(w, "addAnimeResults", p)
		t.ExecuteTemplate(w, "footer", p)

	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	p.Title = "laxMunchies"
	t, _ := template.ParseFiles("web/home.html", "web/header.html", "web/footer.html")
	t.ExecuteTemplate(w, "header", p)
	t.ExecuteTemplate(w, "home", p)
	t.ExecuteTemplate(w, "footer", p)
}
