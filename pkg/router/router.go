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
	request := make(map[string]interface{}) // payload receieved
	message := make(map[string]interface{}) //return message AKA payload

	logger.Log.Infof("Request received from %v", r.RemoteAddr)
	w.Header().Set("Content-type", "application/json") //set header

	err := json.NewDecoder(r.Body).Decode(&request) //decode body and put in var request
	if err != nil {
		logger.Log.Error(err)
	}

	if len(request) > 2 { //Check if more then two key was recieved, if true then drop request
		w.WriteHeader(http.StatusBadRequest)
		message["Error"] = "Given " + strconv.Itoa(len(request)) + " Keys when the max keys available is 2!"
		logger.Log.Errorf("Recieved %v Keys! Max is 2!", len(request))

	} else if _, ok := request["title"]; !ok { //Check if 'title' key is missing, if true then drop request
		w.WriteHeader(http.StatusBadRequest)
		message["Error"] = "'title' key was not found!"
		logger.Log.Error("'title' key was not found in request!")

	} else if request["title"] == "" { //Check if 'title' key is empty, if true then drop request
		w.WriteHeader(http.StatusBadRequest)
		message["Error"] = "'title' key value was Empty!"
		logger.Log.Error("Recivied 'title' key that was Empty!")

	} else { //If all test pass then pass 'title' key value to AniSearch function
		w.WriteHeader(http.StatusOK)
		message = search.AniSearch(request)
	}

	payload, _ = json.MarshalIndent(message, "", "\t") //convert message to JSON format

	logger.Log.Infof("Sending response to %v", r.RemoteAddr)
	w.Write(payload) //Send payload
}
