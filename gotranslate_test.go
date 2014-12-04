package gotranslate

import "testing"

func TestGotranslateNew(t *testing.T) {
	translator, err := New(Japanese, Spanish)

	if err != nil {
		t.Error("An error while getting a translator has ocurred")
	}
	if translator.Translate("こんにちは世界") != "¡Hola Mundo" {
		t.Error("Wrong translation")
	}
}

func TestGotranslateIsLang(t *testing.T) {
	if !IsLang("es") {
		t.Error("IsLang working wrong!")
	}
}

func TestGotranslateQuickTranslation(t *testing.T) {
	if QuickTranslation("hi", English, Spanish) != "hola" {
		t.Error("QuickTranslation working wrong!")
	}
}

func TestGotranslateQueryHistory(t *testing.T) {
	tr, err := New(English, Spanish)
	if err != nil {
		t.Error("Error in New")
	}
	tr.Translate("hi")
	tr.Translate("i love")
	tr.Translate("go")

	if len(tr.QueryHistory()) != 3 {
		t.Error("Wrong length ")
	}
}

func TestGotranslateResultsHistory(t *testing.T) {
	tr, err := New(English, Spanish)
	if err != nil {
		t.Error("Error in New")
	}
	tr.Translate("hi")
	tr.Translate("i love")
	tr.Translate("go")
	if len(tr.ResultsHistory()) != 3 {
		t.Error("Wrong length")
	}
}
