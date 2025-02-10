package main

// source: https://go.dev/doc/tutorial/database-access#single_row
import (
	"database/sql"
	"log"
	"myweatherapi/mysqlAPI/methods/create"
	"myweatherapi/mysqlAPI/methods/delete"
	"myweatherapi/mysqlAPI/methods/read"
	"myweatherapi/mysqlAPI/methods/update"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float32
}

func main() {
	cfg := mysql.Config{
		User:   "mysqlu", //This values were hardcoded to avoid the need of
		Passwd: "pwd",    // exporting variables for each node prior to execute this script
		Net:    "tcp",
		Addr:   "192.168.1.134:3306",
		DBName: "MUSIC",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	r := http.NewServeMux()
	r.HandleFunc("GET /albums", read.GetAlbums(db)) //calls function passing the values for the db connection
	r.HandleFunc("POST /albums/add", create.AddAlbum(db))
	r.HandleFunc("PUT /albums/update", update.UpdateAlbum(db))
	r.HandleFunc("/albums/delete", delete.DeleteAlbum(db))
	log.Print("Starting server on port :8090...")
	log.Fatal(http.ListenAndServe(":8090", r))
}
