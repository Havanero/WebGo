package main

import (
	"database/sql"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"

	gorp "gopkg.in/gorp.v1"
)

type Book struct {
	PK             int64  `db:"pk"`
	Title          string `db:"title"`
	Author         string `db:"author"`
	Classification string `db:"classification"`
	ID             string `db:"id"`
}
type Page struct {
	Books []Book
}
type SearchResult struct {
	Title  string `xml:"title,attr"`
	Author string `xml:"author,attr"`
	Year   string `xml:"hyr,attr"`
	ID     string `xml:"owi,attr"`
}

var db *sql.DB
var dbmap *gorp.DbMap

func initDB() {
	db, _ = sql.Open("sqlite3", "dev.db")
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Book{}, "books").SetKeys(true, "pk")
	dbmap.CreateTablesIfNotExists()
}

func verifyDataBase(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if err := db.Ping(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	next(w, r)
}

func getBookCollections(books *[]Book, sortCol string, w http.ResponseWriter) bool {
	println("running sorting on cols on ", sortCol)
	if sortCol != "title" && sortCol != "author" && sortCol != "classification" {
		sortCol = "pk"
	}
	if _, err := dbmap.Select(books, "select * from books order by "+sortCol); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	return true
}

type ClassifySearchResponse struct {
	Results []SearchResult `xml:"works>work"`
}

type ClassifyBookResponse struct {
	BookData struct {
		Title  string `xml:"title,attr"`
		Author string `xml:"author,attr"`
		ID     string `xml:"owi,attr"`
	} `xml:"work"`
	Classification struct {
		MostPopular string `xml:"sfa,attr"`
	} `xml:"recommendation>ddc>mostPopular"`
}

func find(id string) (ClassifyBookResponse, error) {
	var c ClassifyBookResponse
	body, err := classifyAPI("http://classify.oclc.org/classify2/Classify?&summary=true&owi=" + url.QueryEscape(id))

	if err != nil {
		return ClassifyBookResponse{}, err
	}
	err = xml.Unmarshal(body, &c)
	return c, err
}
func search(query string) ([]SearchResult, error) {
	body, err := classifyAPI("http://classify.oclc.org/classify2/Classify?&summary=true&title=" + url.QueryEscape(query))
	if err != nil {
		return []SearchResult{}, err
	}
	var c ClassifySearchResponse
	err = xml.Unmarshal(body, &c)
	return c.Results, err
}
func classifyAPI(url string) ([]byte, error) {
	var res *http.Response
	var err error
	if res, err = http.Get(url); err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
