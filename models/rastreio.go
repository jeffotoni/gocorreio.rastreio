package models

import (
	"encoding/xml"
)

type Rastreio struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Header  struct {
		Text                   string `xml:",chardata"`
		XOPNETTransactionTrace struct {
			Text                   string `xml:",chardata"`
			XOPNETTransactionTrace string `xml:"X-OPNET-Transaction-Trace,attr"`
		} `xml:"X-OPNET-Transaction-Trace"`
	} `xml:"Header"`
	Body struct {
		Text                 string `xml:",chardata"`
		BuscaEventosResponse struct {
			Text   string `xml:",chardata"`
			Ns2    string `xml:"ns2,attr"`
			Return Return `xml:"return"`
		} `xml:"buscaEventosResponse"`
	} `xml:"Body"`
}

type Return struct {
	Text       string `xml:",chardata"`
	VersionApp string `xml:"versao_app"`
	Versao     string `xml:"versao"`
	Qtd        string `xml:"qtd"`
	Objeto     struct {
		Text      string `xml:",chardata"`
		Numero    string `xml:"numero"`
		Erro      string `xml:"erro"`
		Sigla     string `xml:"sigla"`
		Nome      string `xml:"nome"`
		Categoria string `xml:"categoria"`
		Evento    []struct {
			Text      string `xml:",chardata"`
			Tipo      string `xml:"tipo"`
			Status    string `xml:"status"`
			Data      string `xml:"data"`
			Hora      string `xml:"hora"`
			Descricao string `xml:"descricao"`
			Local     string `xml:"local"`
			Codigo    string `xml:"codigo"`
			Cidade    string `xml:"cidade"`
			Uf        string `xml:"uf"`
		} `xml:"evento"`
	} `xml:"objeto"`
}
