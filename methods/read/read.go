package read

import (
	"database/sql"
	"encoding/json"
	"myweatherapi/mysqlAPI/model"
	"net/http"
)

func GetAlbums(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var albums []model.Album
		rows, err := db.Query("SELECT * FROM Albums")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var alb model.Album
			if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			albums = append(albums, alb)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if len(albums) > 0 {
			if err := json.NewEncoder(w).Encode(albums); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			w.Write([]byte("No albums found"))
		}
	}
}
