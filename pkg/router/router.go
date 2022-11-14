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
	//Get port number from ENV file
	port := ":" + os.Getenv("SERVER_PORT")

	router.HandleFunc("/search", getAnime).Methods("GET")
	logger.Log.Infof("Started server on port %v", port)
	logger.Log.Fatal(http.ListenAndServe(port, router))
}

func getAnime(w http.ResponseWriter, r *http.Request) {

	var payload []byte
	message := make(map[string]interface{})

	//Set Headers
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	logger.Log.Infof("Request received from %v", r.RemoteAddr)

	query := queryStringConverter(r)

	if ok, err := validate(query); ok {

		w.WriteHeader(http.StatusOK)
		//convert message to JSON format
		payload, _ = json.MarshalIndent(search.AniSearch(query), "", "\t")
	} else {
		message["Error"] = err
		//convert message to JSON format
		payload, _ = json.MarshalIndent(message, "", "\t")
		logger.Log.Error(err)
	}

	logger.Log.Infof("Sending response to %v", r.RemoteAddr)

	w.Write(payload) //Send payload

}

func queryStringConverter(r *http.Request) search.Query {
	var q search.Query

	logger.Log.Info(r.URL)

	q.Title = r.URL.Query().Get("title")

	q.Filter = r.URL.Query()["filter"]

	seed, err := strconv.Atoi(r.URL.Query().Get("minSeed"))

	if err == nil {
		q.MinSeed = seed
	}

	entry, err := strconv.Atoi(r.URL.Query().Get("maxEntry"))

	if err == nil {
		q.MaxEntry = entry
	}

	return q

}

func validate(q search.Query) (bool, string) {

	if q.Title == "" {

		return false, "Title is Missing!"
	} else if len(q.Filter) > 10 {

		return false, "You have " + strconv.Itoa(len(q.Filter)) + "! Max is 10!"
	} else if q.MaxEntry > 100 {
		return false, "You have " + strconv.Itoa(q.MaxEntry) + "! Max is 100!"
	} else {
		return true, ""
	}
}
