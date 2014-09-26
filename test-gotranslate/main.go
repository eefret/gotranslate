package main

import (
	"fmt"
	"github.com/eefret/gotranslate"
	"reflect"
)

func main() {
	text, err := gotranslate.TranslationRequest("你好世界", gotranslate.Chinese_T, gotranslate.English)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
