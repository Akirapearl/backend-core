package create

import (
	"database/sql"
	"encoding/json"
	"myweatherapi/mysqlAPI/model"
	"net/http"
)

func AddAlbum(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var alb model.Album
		/*Standardize expected response as a JSON*/
		err := json.NewDecoder(r.Body).Decode(&alb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result, err := db.Exec("INSERT into Albums VALUES ((select max(ID)+1 from Albums a),?, ?, ?);", alb.Title, alb.Artist, alb.Price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		alb.ID = int(id)
		json.NewEncoder(w).Encode(alb)
	}
}

/*
curl http://localhost:8090/albums/add \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"Title": "The last stand","Artist": "Sabaton","Price": 19.99}'

*/
