package router

import (
	"animescrapper/pkg/logger"
	"animescrapper/pkg/search"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()
	port := os.Getenv("SERVER_PORT") //GRAB FROM ENV FILE

	router.HandleFunc("/", getAnime).Methods("GET")
	logger.Log.Infof("Started server on port %v", port)
	logger.Log.Fatal(http.ListenAndServe(port, router))
}

func getAnime(w http.ResponseWriter, r *http.Request) {

	var payload []byte
	var request search.Query                // Create Query Struct
	message := make(map[string]interface{}) //return message AKA payload

	logger.Log.Infof("Request received from %v", r.RemoteAddr)
	w.Header().Set("Content-type", "application/json") //set header

	err := json.NewDecoder(r.Body).Decode(&request) //decode body and put in var request
	if err != nil {
		logger.Log.Error(err)
	}

	if ok, err := validate(request); ok {
		w.WriteHeader(http.StatusOK)
		message = search.AniSearch(request)

		logger.Log.Info(request)
	} else {
		message["Error"] = err
	}

	payload, _ = json.MarshalIndent(message, "", "\t") //convert message to JSON format

	logger.Log.Infof("Sending response to %v", r.RemoteAddr)
	w.Write(payload) //Send payload
}

func validate(q search.Query) (bool, string) {

	if q.Title == "" {

		return false, "Title is Missing!"
	} else if len(q.Filter) > 10 {

		return false, "You have " + strconv.Itoa(len(q.Filter)) + "! Max is 10!"
	} else {
		return true, ""
	}
}
