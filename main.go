package main

import (
	route "./route"
	"html/template"
	"net/http"
)

func serverweb(w http.ResponseWriter, _ *http.Request) {
	file := "index.html"
	tmpl := template.New("index.html")
	tmpl = template.Must(tmpl.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "index.html", file)
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", serverweb)
	http.ListenAndServe(":8080", log(server))
}

func log(getlogerror http.Handler) http.Handler {
	a := route.Route(getlogerror)
	return http.HandlerFunc(a)

}
