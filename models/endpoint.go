package models

type end struct {
	Method string
	Source string
	Url    string
	Body   string
}

var Endpoints = []end{
	{"POST", "correio", "http://webservice.correios.com.br:80/service/rastro",
		`<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
		    <Body>
		        <buscaEventos xmlns="http://resource.webservice.correios.com.br/">
		            <usuario xmlns="">%s</usuario>
		            <senha xmlns="">%s</senha>
		            <tipo xmlns="">%s</tipo>
		            <resultado xmlns="">%s</resultado>
		            <lingua xmlns="">101</lingua>
		            <objetos xmlns="">%s</objetos>
		        </buscaEventos>
		    </Body>
		</Envelope>
`}}
