package restful

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	Message string `json:"message"`
}

func HandleGet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hi! Get işlemi başarılı")
}

func HandlePost(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var requestData Data
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("post isteği:", requestData.Message)
	response := Data{Message: "POST işlemi başarılı"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("application/json", "application/json")
	w.Write(jsonResponse)
}
