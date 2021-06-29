package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
	a := func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.RemoteAddr, "\t", r.RequestURI, r.Method)
		getlogerror.ServeHTTP(w, r)
		fmt.Println(r.FormValue("num1"), r.FormValue("num2"), r.FormValue("operator"))


		num1, _ := strconv.ParseFloat(r.FormValue("num1"),64)
		num2, _ := strconv.ParseFloat(r.FormValue("num2"),64)

		var Res float64

		switch r.FormValue("operator") {
		case "add":
			Res = num1 + num2
		case "sub":
			Res = num1 - num2
		case "mult":
			Res = num1 * num2
		case "div":
			Res = num1 / num2
		}

		fmt.Println(Res)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		_, _ = fmt.Fprint(w, "your result : <br>"+fmt.Sprintf("%f", Res))

	}
	return http.HandlerFunc(a)

}


