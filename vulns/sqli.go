package vulns

import (
	"log"
	"io"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


func Sqli(w http.ResponseWriter, r *http.Request) {
	const file string = "activities.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var id int
	var date string
	var description string
	err = db.QueryRow("SELECT * FROM activities WHERE id=" + r.URL.Query().Get("x")).Scan(&id,&date, &description)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, description)

}