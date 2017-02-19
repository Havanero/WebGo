package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	//sessions "github.com/goincremental/negroni-sessions"
	gmux "github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func rootPage(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	templates := template.Must(template.ParseFiles("templates/base.html", "templates/web-go/index.html"))
	/*var sortColumn string
	  if sortBy := sessions.GetSession(r).Get("sortBy"); sortBy != nil {
	    sortColumn = sortBy.(string)
	  }*/
	p := Page{Books: []Book{}}
	if !getBookCollections(&p.Books, r.FormValue("sortBy"), w) {
		return
	}
	if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	session.Values["foo"] = "bar"
	session.Save(r, w)
}
func getBooks(w http.ResponseWriter, r *http.Request) {
	var b []Book
	if !getBookCollections(&b, r.FormValue("sortBy"), w) {
		return
	}
	//	sessions.GetSession(r).Set("SortBy", r.FormValue("sortBy"))
	if err := json.NewEncoder(w).Encode(b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func postSearch(w http.ResponseWriter, r *http.Request) {
	var results []SearchResult
	var err error
	if results, err = search(r.FormValue("search")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func addBooks(w http.ResponseWriter, r *http.Request) {
	var book ClassifyBookResponse
	var err error

	if book, err = find(r.FormValue("id")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	b := Book{
		PK:             -1,
		Title:          book.BookData.Title,
		Author:         book.BookData.Author,
		Classification: book.Classification.MostPopular,
	}
	if err = dbmap.Insert(&b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	pk, _ := strconv.ParseInt(gmux.Vars(r)["pk"], 10, 64)
	if _, err := dbmap.Delete(&Book{pk, "", "", "", ""}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
