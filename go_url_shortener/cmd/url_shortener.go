package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/thaiminh2022/go_url_shortener/internal"
)

var tpl *template.Template

func init() {
	localTpl, err := template.ParseGlob("web/*.html")
	tpl = localTpl

	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	mux := http.NewServeMux()
	//middlewares := middleware.CreateStack(
	//	middleware.Logging,
	//)

	internal.Handler(mux, tpl)
	// Serve uup some css

	fmt.Println("Listening at port 3000")
	http.ListenAndServe(":3000", mux)
}
