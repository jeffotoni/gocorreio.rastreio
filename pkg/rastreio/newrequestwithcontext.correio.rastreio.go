package rastreio

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/jeffotoni/gocorreio.rastreio/models"
	"net/http"
)

func NewRequestWithContextCorreioRastreio(ctx context.Context, cancel context.CancelFunc, rastreio, source, method, endpoint,
	payload string, chResult chan<- Result) {

	var err error
	payload = fmt.Sprintf(payload, rastreio)
	req, err := http.NewRequestWithContext(ctx, method, endpoint, bytes.NewReader([]byte(payload)))
	if err != nil {
		return
	}

	req.Header.Set("Content-type", "text/xml; charset=utf-8")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	obj := new(models.Rastreio)
	err = xml.NewDecoder(response.Body).Decode(obj)
	if err == nil {
		c := obj.Body.BuscaEventosResponse.Return

		b, err := json.Marshal(c)
		if err == nil {
			chResult <- Result{Body: b}
			cancel()
		}
	}
	return
}
