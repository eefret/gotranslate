package main

import (
	"fmt"
	"github.com/eefret/gotranslate"
	//"reflect"
)

func main() {
	text, err := gotranslate.TranslationRequest("你好世界", gotranslate.AUTOMATIC, "es")
	check(err)
	fmt.Println(text)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
