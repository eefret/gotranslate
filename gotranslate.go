/*
Copyright 2014 Kaissersoft Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package gotranslate

import (
	// "errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	// "reflect"
	"regexp"
	"strings"
)

const urlFormat string = "http://translate.google.com/translate_a/t?client=t&text=%s&hl=en&sl=%s&tl=%s&ie=UTF-8&oe=UTF-8&multires=1&otf=1&pc=1&trs=1&ssel=3&tsel=6&sc=1"

/**
*Static method for translating text
*@param string text Text to translate
*@param Lang from Language from which will be translated
*@param Lang to Language to which will be translated
 */
func QuickTranslation(text string, from Lang, to Lang) (string, error) {

	return "", nil
}

type Response [][][]string

func TranslationRequest(text string, from Lang, to Lang) (string, error) {

	/*	if !checkLang(from) && !checkLang(to) {
		return "", errors.New("Not Lang type")
	}*/

	var Url *url.URL
	Url, err := url.Parse("http://translate.google.com")
	check(err)

	text = strings.Replace(text, "\"", "", -1)
	text = strings.Replace(text, "[", "", -1)
	text = strings.Replace(text, "]", "", -1)
	text = strings.Replace(text, ",", " ", -1)

	Url.Path += "/translate_a/t"
	parameters := url.Values{}
	parameters.Add("client", "t")
	parameters.Add("text", text)
	parameters.Add("hl", "en")
	parameters.Add("sl", string(from))
	parameters.Add("tl", string(to))
	parameters.Add("ie", "UTF-8")
	parameters.Add("oe", "UTF-8")
	parameters.Add("multires", "1")
	parameters.Add("otf", "1")
	parameters.Add("pc", "1")
	parameters.Add("trs", "1")
	parameters.Add("ssel", "3")
	parameters.Add("tsel", "6")
	parameters.Add("sc", "1")
	Url.RawQuery = parameters.Encode()

	resp, err := http.Get(Url.String())
	defer resp.Body.Close()
	check(err)

	contents, err := ioutil.ReadAll(resp.Body)
	check(err)

	reg, err := regexp.Compile("\"(.+?)\"")
	check(err)

	var allStrings []string
	allStrings = reg.FindAllString(string(contents), 3)

	return allStrings[0], nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*func checkLang(i interface{}) bool {
	if i.type == Lang {
		return true
	}
	return false
}*/
