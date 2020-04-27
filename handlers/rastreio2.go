package handler

import (
	"github.com/jeffotoni/gocorreio.rastreio/pkg/rastreio"
	"net/http"
	"net/url"
)

func Rastreio2(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	endpoint := r.URL.Path
	if endpoint != "/api/v2/" {
		w.WriteHeader(http.StatusFound)
		return
	}

	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, "Codigo de rastreio invalido e campos invalidos", http.StatusBadRequest)
		return
	}
	q := u.Query()
	var usuario, senha, etiqueta string
	var tipo string = "L"
	var resultado string = "T"

	if len(q.Get("usuario")) <= 0 {
		http.Error(w, "Campo usuario é obrigatorio", http.StatusBadRequest)
		return
	}
	if len(q.Get("senha")) <= 0 {
		http.Error(w, "Campo senha é obrigatorio", http.StatusBadRequest)
		return
	}
	if len(q.Get("etiqueta")) <= 0 {
		http.Error(w, "Campo etiqueta é obrigatorio", http.StatusBadRequest)
		return
	}
	if len(q.Get("tipo")) > 0 {
		tipo = q.Get("tipo")
	}
	if len(q.Get("resultado")) <= 0 {
		resultado = q.Get("resultado")
	}

	etiqueta = q.Get("etiqueta")
	usuario = q.Get("usuario")
	senha = q.Get("senha")

	result, err := rastreio.Search2(usuario, senha, etiqueta, tipo, resultado)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(result))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
	return
}
