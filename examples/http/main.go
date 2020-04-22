package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/jeffotoni/gocorreio.rastreio/pkg/rastreio"
)

var (
	Port = ":8080"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/etiqueta/", Handlerrastreio)
	mux.HandleFunc("/etiqueta", NotFound)
	mux.HandleFunc("/", NotFound)

	server := &http.Server{
		Addr:    Port,
		Handler: mux,
	}

	log.Println("port", Port)
	log.Fatal(server.ListenAndServe())
}

func Handlerrastreio(w http.ResponseWriter, r *http.Request) {

	rastreiostr := strings.Split(r.URL.Path[1:], "/")[1]
	if len(rastreiostr) != 8 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := rastreio.Search(rastreiostr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(result))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
	return
}
