package route

import (
	"fmt"
	"net/http"
	"strconv"
)

func Route(log http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.RemoteAddr, "\t", r.RequestURI, r.Method)
		log.ServeHTTP(w, r)
		fmt.Println(r.FormValue("num1"), r.FormValue("num2"), r.FormValue("operator"))

		num1, _ := strconv.ParseFloat(r.FormValue("num1"), 64)
		num2, _ := strconv.ParseFloat(r.FormValue("num2"), 64)

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
}
