package handler

import (
	"github.com/jeffotoni/gocorreio.rastreio/pkg/rastreio"
	"net/http"
	"strings"
)

func Rastreio(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	validEndpoint := strings.Split(r.URL.Path, "/")
	if len(validEndpoint) > 4 {
		w.WriteHeader(http.StatusFound)
		return
	}

	etiqueta := strings.Split(r.URL.Path[2:], "/")[2]
	if len(etiqueta) <= 0 {
		http.Error(w, "Codigo de rastreio invalido", http.StatusBadRequest)
		return
	}

	result, err := rastreio.Search(etiqueta)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(result))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
	return
}
