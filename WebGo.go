package main

import (
	"flag"
	"net/http"

	"github.com/codegangsta/negroni"
	gmux "github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	initDB()
	var dir string
	flag.StringVar(&dir, "dir", "assets", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	mux := gmux.NewRouter()
	mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(dir))))

	mux.HandleFunc("/", rootPage).Methods("GET")
	mux.HandleFunc("/books", getBooks).Methods("GET")
	mux.HandleFunc("/search", postSearch).Methods("POST")
	mux.HandleFunc("/books", addBooks).Methods("PUT")
	mux.HandleFunc("/books/{pk}", deleteBooks).Methods("DELETE")

	n := negroni.Classic()

	//store := cookiestore.New([]byte("secret123"))
	//n.Use(sessions.Sessions("my_session", store))

	n.Use(negroni.HandlerFunc(verifyDataBase))
	n.UseHandler(mux)
	n.Run(":8081")
}
