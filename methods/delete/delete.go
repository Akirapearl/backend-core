package delete

import (
	"database/sql"
	"encoding/json"
	"myweatherapi/mysqlAPI/model"
	"net/http"
)

func DeleteAlbum(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var alb model.Album
		err := json.NewDecoder(r.Body).Decode(&alb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = db.Exec("DELETE FROM Albums WHERE ID = ?", alb.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

/*
  -- REQUIRES TO AVOID ESCAPING LINES
	curl -X DELETE http://localhost:8090/albums/delete -H "Content-Type: application/json" -d '{"ID" : 8}'
*/
