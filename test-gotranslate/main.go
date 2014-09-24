package main

import (
	"fmt"
	"github.com/eefret/gotranslate"
)

func main() {
	fmt.Println(gotranslate.TranslationRequest("hola \" mmg\" mundo", gotranslate.ES, gotranslate.FR))
}
