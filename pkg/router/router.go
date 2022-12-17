package router

import (
	"animescrapper/pkg/logger"
	"animescrapper/pkg/search"

	//"animescrapper/pkg/search"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()
	//Get port number from ENV file
	port := ":" + os.Getenv("SERVER_PORT")

	router.HandleFunc("/search", getAnimeTorrents).Methods("GET")

	logger.Log.Infof("Started server on port %v", port)
	logger.Log.Fatal(http.ListenAndServe(port, router))
}

func getAnimeTorrents(w http.ResponseWriter, r *http.Request) {

	var payload []byte
	//message := make(map[string]interface{})

	//Log Connection
	logger.Log.Infof("Request received from %v", r.RemoteAddr)

	//Set Headers
	setHeaders(w)

	query, err := queryStringConverter(r)

	if err != "" {
		sendError(w, err)
		return
	}

	if ok, err := validate(query); ok {

		w.WriteHeader(http.StatusOK)
		//convert message to JSON format
		//payload, _ = json.MarshalIndent(message, "", "\t")
		payload, _ = json.MarshalIndent(search.AniSearch(query), "", "\t")
	} else {
		sendError(w, err)
	}

	logger.Log.Infof("Sending response to %v", r.RemoteAddr)

	w.Write(payload) //Send payload
}
