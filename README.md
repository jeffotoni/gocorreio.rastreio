# gocorreio.rastreio

Um simples pacote para buscar nos correios os eventos de postagem, o rastreamento de pedidos postados do seu pedido.

Podendo implementar para ter uma saída ainda mais completa conforme sua necessidade, então fique a vontade em alterar conforme seu cenário.

O server é extremamente rápido, e usa cache em memória ele está configurado para 2G de Ram, caso queira alterar está tudo bonitinho no /config.

gocorreio.rastreio também poderá ser usado como Lib, ou seja você irá conseguir fazer um import em seu pkg/rastreio  e fazer a chamada direto do seu método em seu código.

## Usar como Lib Search v1
```go

package main

import (
	"fmt"
	"github.com/jeffotoni/gocorreio.rastreio/pkg/rastreio"
)

func main() {

	result, err := rastreio.Search("PX521577722BR")
	fmt.Println(err)
	fmt.Println(result)
}

```

## Usar como Lib Search v2
```go

package main

import (
   "fmt"
   "github.com/jeffotoni/gocorreio.rastreio/pkg/rastreio"
)

func main() {

   //usuario, senha, etiqueta, tipo e resultado
   result, err := rastreio.Search2("ECT","SRO","PX521577722BR","L", "T")
   fmt.Println(err)
   fmt.Println(result)
}

```

Ou se preferir for criar seu próprio serviço e sua api basta fazer como exemplo abaixo:
Existe em examples dois exemplos de commo integrar a lib gocorreio.rastreio em seu projeto.

```bash

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/etiqueta/", func(w http.ResponseWriter, r *http.Request){
		etiquetaStr := strings.Split(r.URL.Path[1:], "/")[1]
		if len(etiquetaStr) != 8 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		result, err := rastreio.Search(etiquetaStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	})

	log.Fatal(http.ListenAndServe(":8080"))
}

```

Você pode fazer seu próprio build usando Go, ou você poderá utilizar docker-compose. O server irá funcionar na porta 8085, mas caso queira alterar basta ir na pasta /config.

Para subir o serviço para seu Servidor ou sua máquina local basta compilar, e a porta 8085 será aberta para consumir o endpoint /api/v1/{etiqueta}

# Install gocorreio.rastreio

Caso queira utilizar ele como serviço, basta baixa-lo ou usar o docker para utilizado.

## linux bash
```bash
$ git clone https://github.com/jeffotoni/gocorreio.rastreio
$ cd gocorreio.rastreio
$ go build -ldflags="-s -w" 
$ ./gocorreio.rastreio
$ 2020/04/21 12:56:46 Port: :8085

```

## docker e docker-compose

Deixei um script para facilitar a criação de sua imagem, todos os arquivos estão na raiz, docker-compose.yaml, Dockerfile tudo que precisa para personalizar ainda mais se precisar.
Ao rodar o script ele irá fazer pull da imagem que encontra-se no hub.docker.
```bash

$ sh deploy.gocorreio.rastreio.sh

```

## Listando service
```bash
$ docker-compose ps
Creating gocorreio.rastreio ... done
Name    Command   State           Ports         
------------------------------------------------
gocorreio.rastreio   /gocorreio.rastreio    Up      0.0.0.0:8085->8085/tcp
-e Generated Run docker-compose [ok] 

```

## Executando sua API V1
```bash

$ curl -i http://localhost:8085/api/v1/PX521577722BR

```


## Executando sua API V2
```bash

$ curl -i -XGET "http://localhost:8085/api/v2/?usuario=ECT&senha=SRO&etiqueta=PX521577722BR&tipo=L&resultado=T"

```

## Saida Json (erro)
```json

{
   "Text":"",
   "Versao":"2.0",
   "Qtd":"1",
   "Objeto":
   {
      "Text":"",
      "Numero":"PX521577733BC",
      "Erro":"Objeto não encontrado na base de dados dos Correios.",
      "Sigla":"",
      "Nome":"",
      "Categoria":"",
      "Evento":null
   }
}

```

## Saida Json (sucesso)

```json

{
   "Text":"",
   "Versao":"2.0",
   "Qtd":"1",
   "Objeto":{
      "Text":"",
      "Numero":"PX521577722BR",
      "Erro":"",
      "Sigla":"PX",
      "Nome":"ENCOMENDA PAC (ETIQUETA LOGICA) ",
      "Categoria":"ENCOMENDA PAC",
      "Evento":[
         {
            "Text":"",
            "Tipo":"PO",
            "Status":"01",
            "Data":"29/11/2019",
            "Hora":"17:31",
            "Descricao":"Objeto postado",
            "Local":"AGF DOUTOR JOAO MENDES",
            "Codigo":"01502994",
            "Cidade":"SAO PAULO",
            "Uf":"SP"
         }
      ]
   }
}

```

