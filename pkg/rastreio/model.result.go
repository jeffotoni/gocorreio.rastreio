package rastreio

import "github.com/jeffotoni/gocorreio.rastreio/models"

type Result struct {
	//Rastreio models.Rastreio
	Body []byte
}

var chResult = make(chan Result, len(models.Endpoints))
