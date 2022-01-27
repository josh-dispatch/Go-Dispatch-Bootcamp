package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String("listen-addr", ":8080", "HTTP address to listen to")
)

func main() {
	flag.Parse()

	router := http.NewServeMux()

	router.HandleFunc("/data", handleMethod(http.MethodGet, data))

	if err := http.ListenAndServe(*addr, router); err != nil {
		log.Fatal(err)
	}
}

func handleMethod(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		handler.ServeHTTP(w, r)
	}
}

func data(w http.ResponseWriter, r *http.Request) {

	data, csvErr := GetCsvData()

	buf, err := json.Marshal(map[string]interface{}{
		"data": data,
	})
	if err != nil || csvErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(buf)
}
