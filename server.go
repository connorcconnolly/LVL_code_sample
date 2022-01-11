//server.go
//Connor Connolly
//January 9th 2022
package main

import(
  _ "github.com/mattn/go-sqlite3"
  "fmt"
  "log"
  "net/http"
  "database/sql"
  "strings"
)
type Track struct {
  Name string `json:"Name"`
  Composer string `json:"Composer"`
  Album string `json:"Album"`
}

func main(){
  requestHandler()
}

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Println("Listening on port 4041")
}

func search(w http.ResponseWriter, r *http.Request){
  params, ok := r.URL.Query()["param"]

  if !ok || len(params[0]) < 1 {
    log.Println("Missing search term")
    return
  }
  searchTerm:=params[0]

  db, err := sql.Open("sqlite3","./Chinook_Sqlite.sqlite")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()
  searchTerm = strings.TrimSuffix(searchTerm, "\n")

  //Query db using search term
  rows, err := db.Query("SELECT Name FROM Track WHERE Name LIKE '%"+searchTerm+"%'")
  if err != nil {
		log.Fatal(err)
	}
  defer rows.Close()

	err = rows.Err()
  if err != nil {
		log.Fatal(err)
	}
  fmt.Println("Tracks Matching the term: ",searchTerm)

  //output to stdout
  trackData := Track{}

  for rows.Next(){

    err = rows.Scan(&trackData.Name)
    if err != nil {
  		log.Fatal(err)
  	}
    fmt.Println(trackData.Name)
  }
}

func requestHandler() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/search", search)
    log.Fatal(http.ListenAndServe(":4041", nil))

}
