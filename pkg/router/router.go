package router

import (
	"animescrapper/pkg/logger"
	"animescrapper/pkg/search"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()
	port := 8080

	router.HandleFunc("/", getAnime).Methods("GET")
	logger.Log.Infof("Started server on port %v", port)
	logger.Log.Fatal(http.ListenAndServe(":8080", router))
}

func getAnime(w http.ResponseWriter, r *http.Request) {

	var payload []byte
	request := make(map[string]interface{})
	message := make(map[string]interface{})

	logger.Log.Infof("Request received from %v", r.RemoteAddr)
	w.Header().Set("Content-type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Log.Error(err)
	}

	if len(request) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		message["Error"] = "Given " + strconv.Itoa(len(request)) + " Keys when only 1 was Expected!"
		logger.Log.Errorf("Recieved %v Keys instead of one! Dropping...", len(request))

	} else if _, ok := request["title"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		message["Error"] = "'title' key was not found!"
		logger.Log.Error("title key was not found in request! Dropping...")

	} else if request["title"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		message["Error"] = "'title' key value was Empty!"
		logger.Log.Error("Recivied title key that was Empty! Dropping...")

	} else {
		w.WriteHeader(http.StatusOK)
		message = search.AniSearch(request["title"].(string))
		logger.Log.Info(payload)
	}

	payload, _ = json.MarshalIndent(message, "", "\t")

	logger.Log.Infof("Sending response to %v", r.RemoteAddr)
	w.Write(payload)
}
