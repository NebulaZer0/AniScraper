package router

import (
	"animescrapper/pkg/logger"
	"animescrapper/pkg/search"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ACAO"))
}

func queryStringConverter(r *http.Request) (search.Query, string) {
	var q search.Query

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

	return q, ""

}

func validate(q search.Query) (bool, string) {
	// logger.Log.Info(len(q.Filter))
	if q.Title == "" {

		return false, "Title is Missing!"
	} else if len(q.Filter) > 10 {

		return false, "filter has" + strconv.Itoa(len(q.Filter)) + " values! Max is 10!"
	} else if q.MaxEntry > 100 {
		return false, "maxEntry value is " + strconv.Itoa(q.MaxEntry) + "! Max is 100!"
	} else {
		return true, ""
	}

}

func sendError(w http.ResponseWriter, err string) {
	var payload []byte
	logger.Log.Error(err)
	message := make(map[string]string)

	w.WriteHeader(http.StatusBadRequest)

	message["error"] = err

	payload, _ = json.Marshal(message)

	w.Write(payload)
}
