package internal

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/thaiminh2022/go_url_shortener/internal/tools"
)

func Handler(mux *http.ServeMux, tpl *template.Template) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.html", nil)
	})

	mux.HandleFunc("/r/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := strings.Trim(r.PathValue("id"), " ")
		if id == "" {
			tpl.ExecuteTemplate(w, "400.html", "cannot find url id, please check your url again")
			return
		}

		var db *tools.DatabaseInterface
		db, err := tools.NewDatabase()

		if err != nil {
			log.Println(err)
			tpl.ExecuteTemplate(w, "500.html", err)
			return
		}
		url, err := (*db).GetUrlAtKey(id)
		if err != nil {
			log.Println(err)
			tpl.ExecuteTemplate(w, "400.html", err)
			return
		}

		http.Redirect(w, r, url, http.StatusPermanentRedirect)

	})
	mux.HandleFunc("/s", func(w http.ResponseWriter, r *http.Request) {
		url := strings.Trim(r.PostFormValue("url"), " ")
		if url == "" {
			tpl.ExecuteTemplate(w, "400.html", "cannot find url to shorten, are you sure you submit the form correctly?")
			log.Println("url is empty")
			return
		}
		log.Println("URL is:", url)
		urlID := GetRandomID()

		// Store database with url and url id
		var db *tools.DatabaseInterface
		db, err := tools.NewDatabase()
		if err != nil {
			log.Println(err)
			tpl.ExecuteTemplate(w, "500.html", err)
			return
		}

		key, existed := (*db).UrlExists(url)
		if existed {
			tpl.ExecuteTemplate(w, "shorten.html", key)
			return
		}

		err = (*db).StoreKeyUrl(urlID, url)
		for err != nil {
			urlID = GetRandomID()
			err = (*db).StoreKeyUrl(urlID, url)
		}

		err = (*db).SaveDB()
		if err != nil {
			log.Println(err)
			tpl.ExecuteTemplate(w, "500.html", err)
			return
		}
		tpl.ExecuteTemplate(w, "shorten.html", urlID)
	})
}
