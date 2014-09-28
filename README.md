GoTranslate 
=======

Author: Christopher T. Herrera (eefretsoul AT gmail DOT com)

This is a free google translate api for golang, it uses google translate. 

Project Usage
-------------
The API usage is very simple. Just import the gotranslate package

	import (
		"github.com/eefret/gotranslate"
	)

You can use the quick translation:
	
	//translate "Hola" from Spanish (ES) to English(EN)
	gotranslate.QuickTranslation("Hola",gotranslate.Spanish, gotranslate.EN)

Or create a Translator struct, to ease a lot of translations and save a history:
	
	//Passed source lang, target lang and historySize
	t, err := gotranslate.GetTranslator(gotranslate.ES, gotranslate.EN, 20)
	t.Translate("hola") //prints "hello"
	t.Translate("mundo") //prints "world"
	querySlice := t.GetQueryHistory() // a slice with ["hola","mundo"]
	resultSlice := t.GetResultsHistory() // a slice with ["hello", "world"]


See the project documentation to see all Available Langs and other Stuff.

Project Documentation
---------------------
The automatically generated documentation can be found in godocs.
[![GoDoc](https://godoc.org/github.com/eefret/gotranslate?status.svg)](https://godoc.org/github.com/eefret/gotranslate)
