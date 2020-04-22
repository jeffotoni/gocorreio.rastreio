package codigoRastreio

import (
	"context"
	"github.com/jeffotoni/gocodigoRastreio/config"
	"github.com/jeffotoni/gocodigoRastreio/models"
	"github.com/jeffotoni/gocodigoRastreio/service/ristretto"
	"time"
)

type Result struct {
	Body []byte
}

var chResult = make(chan Result, len(models.Endpoints))

func Search(codigoRastreio string) (string, error) {

	jsoncodigoRastreio := ristretto.Get(codigoRastreio)
	if len(jsoncodigoRastreio) > 0 {
		return jsoncodigoRastreio, nil
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, e := range models.Endpoints {
		endpoint := e.Url
		source := e.Source
		method := e.Method
		payload := e.Body
		go func(cancel context.CancelFunc, codigoRastreio, method, source, endpoint, payload string, chResult chan<- Result) {

			NewRequestWithContextCorreioRastreio(ctx, cancel, codigoRastreio, source, method, endpoint, payload, chResult)

		}(cancel, codigoRastreio, method, source, endpoint, payload, chResult)
	}

	select {
	case result := <-chResult:
		ristretto.Set(codigoRastreio, string(result.Body))
		return string(result.Body), nil

	case <-time.After(time.Duration(5) * time.Second):
		cancel()
	}

	return config.JsonDefault, nil
}
