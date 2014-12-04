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
/*
	This package is to help all gophers out there to translate texts with
	google translate for free, hope it helps a lot of people
*/
package gotranslate

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

//=======================================================================
//							Const
//=======================================================================

const baseURL string = "http://translate.google.com"

//=======================================================================
//							Errors
//=======================================================================

var (
	//ErrInvalidLang is launched when an invalid Lang is passed to any func
	ErrInvalidLang = errors.New("the lang is invalid, please use a valid one")
	//ErrNoTranslation is launched when there's no translation available
	ErrNoTranslation = errors.New("theres no translation")
)

//=======================================================================
//						Structs & Types
//=======================================================================

//Translator is a struct that contains an origin Lang and a result Lang
//and saves history of the queries made
type Translator struct {
	fromLang      Lang
	toLang        Lang
	queryGroup    []string
	responseGroup []string
}

//=======================================================================
//							Funcs
//=======================================================================

//New returns a Translator struct to ease translation operations
//need a source Lang and a target Lang
func New(from Lang, to Lang) (*Translator, error) {
	if !from.valid() || !to.valid() {
		return nil, ErrInvalidLang
	}

	t := &Translator{
		fromLang:      from,
		toLang:        to,
		queryGroup:    make([]string, 0, 20),
		responseGroup: make([]string, 0, 20),
	}
	return t, nil
}

//Translate takes a string and make the translation over the created Struct
func (t *Translator) Translate(text string) string {
	t.queryGroup = append(t.queryGroup, text)
	txt, err := translationRequest(text, t.fromLang, t.toLang)
	check(err)
	t.responseGroup = append(t.responseGroup, txt)
	return txt
}

//QueryHistory returns all strings submitted to Translator
func (t *Translator) QueryHistory() []string {
	strings := append([]string(nil), t.queryGroup...)
	return strings
}

//ResultsHistory returns all strings obtained from Translator.Translate
func (t *Translator) ResultsHistory() []string {
	strings := append([]string(nil), t.responseGroup...)
	return strings
}

//QuickTranslation translate a single string given from and to langs
func QuickTranslation(text string, from Lang, to Lang) string {
	traslatedText, err := translationRequest(text, from, to)
	check(err)
	return traslatedText
}

func translationRequest(text string, from Lang, to Lang) (string, error) {

	if !from.valid() || !to.valid() {
		return "", ErrInvalidLang
	}

	var URL *url.URL
	URL, err := url.Parse(baseURL)
	check(err)

	text = strings.Replace(text, "\"", "", -1)
	text = strings.Replace(text, "[", "", -1)
	text = strings.Replace(text, "]", "", -1)
	text = strings.Replace(text, ",", " ", -1)

	URL.Path += "/translate_a/t"
	parameters := url.Values{}
	parameters.Add("client", "t")
	parameters.Add("text", text)
	parameters.Add("hl", "en")
	parameters.Add("sl", from.String())
	parameters.Add("tl", to.String())
	parameters.Add("ie", "UTF-8")
	parameters.Add("oe", "UTF-8")
	parameters.Add("multires", "1")
	parameters.Add("otf", "1")
	parameters.Add("pc", "1")
	parameters.Add("trs", "1")
	parameters.Add("ssel", "3")
	parameters.Add("tsel", "6")
	parameters.Add("sc", "1")
	URL.RawQuery = parameters.Encode()

	resp, err := http.Get(URL.String())
	defer resp.Body.Close()
	check(err)

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(http.StatusText(resp.StatusCode))
	}

	contents, err := ioutil.ReadAll(resp.Body)
	check(err)
	reg, err := regexp.Compile("\"(.+?)\"")
	check(err)

	var allStrings []string
	allStrings = reg.FindAllString(string(contents), 2)

	if len(allStrings) < 1 {
		return "", ErrNoTranslation
	}

	s := allStrings[0]
	s = strings.Trim(s, "\"")
	return s, nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
