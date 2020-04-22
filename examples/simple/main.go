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
