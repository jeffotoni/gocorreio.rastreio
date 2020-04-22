package models

type end struct {
	Method string
	Source string
	Url    string
	Body   string
}

var Endpoints = []end{
	{"POST", "correio", "",
		`<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
		    <Body>
		        <buscaEventos xmlns="http://resource.webservice.correios.com.br/">
		            <usuario xmlns="">ECT</usuario>
		            <senha xmlns="">SRO</senha>
		            <tipo xmlns="">L</tipo>
		            <resultado xmlns="">T</resultado>
		            <lingua xmlns="">101</lingua>
		            <objetos xmlns="">%s</objetos>
		        </buscaEventos>
		    </Body>
		</Envelope>
`}}
