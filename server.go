//go:build !js

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func wasmHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("www/index.html"))

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		err := tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
		}
	})
}

func main() {
	fs := http.StripPrefix("/www/", http.FileServer(http.Dir("./www")))
	http.Handle("/www/", fs)

	http.Handle("/home", wasmHandler())
	http.ListenAndServe(":8080", nil)

}

/* OR
import (
    "log"
    "net/http"
    "strings"
)

const dir = "./www"

func main() {
    fs := http.FileServer(http.Dir(dir))
    log.Print("Serving " + dir + " on http://localhost:8080")
    http.ListenAndServe(":8080", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Add("Cache-Control", "no-cache")
        if strings.HasSuffix(req.URL.Path, ".wasm") {
            resp.Header().Set("content-type", "application/wasm")
        }
        fs.ServeHTTP(resp, req)
    }))
}
*/

// go run server.go
