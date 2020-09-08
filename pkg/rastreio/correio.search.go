package rastreio

import (
	"context"
	"runtime"
	"time"

	"github.com/jeffotoni/gocorreio.rastreio/config"
	"github.com/jeffotoni/gocorreio.rastreio/models"
	"github.com/jeffotoni/gocorreio.rastreio/service/ristretto"
)

func Search(codigoRastreio string) (string, error) {
	var chResult = make(chan Result, len(models.Endpoints))
	jsoncodigoRastreio := ristretto.Get(codigoRastreio)
	if len(jsoncodigoRastreio) > 0 {
		return jsoncodigoRastreio, nil
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var usuario string = "ECT"
	var senha string = "SRO"
	var tipo string = "L"
	var resultado string = "T"

	for _, e := range models.Endpoints {
		endpoint := e.Url
		source := e.Source
		method := e.Method
		payload := e.Body
		go func(ctx context.Context, cancel context.CancelFunc, usuario, senha, codigoRastreio, tipo, resultado, method, source, endpoint, payload string, chResult chan<- Result) {

			NewRequestWithContextCorreioRastreio(ctx, cancel, usuario, senha, codigoRastreio, tipo, resultado, source, method, endpoint, payload, chResult)

		}(ctx, cancel, usuario, senha, codigoRastreio, tipo, resultado, method, source, endpoint, payload, chResult)
	}

	select {
	case result := <-chResult:
		ristretto.SetTTL(codigoRastreio, string(result.Body), time.Duration(time.Minute*5))
		return string(result.Body), nil

	case <-time.After(time.Duration(5) * time.Second):
		cancel()
	}

	return config.JsonDefault, nil
}
