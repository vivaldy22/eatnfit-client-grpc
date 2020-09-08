package respJson

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(resp interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
