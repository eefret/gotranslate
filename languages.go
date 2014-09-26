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

//=======================================================================
//							Const
//=======================================================================

const (
	//Afrikaans
	Afrikaans Lang = "af"
	//Afrikaans
	AF Lang = "af"
	//Albanian
	Albanian Lang = "sq"
	//Albanian
	SQ Lang = "sq"
	//Arabic
	Arabic Lang = "ar"
	//Arabic
	AR Lang = "ar"
	//Armenian
	Armenian Lang = "hy"
	//Armenian
	HY Lang = "hy"
	//Azerbaijani
	Azerbaijani Lang = "az"
	//Azerbaijani
	AZ Lang = "az"
	//Basque
	Basque Lang = "eu"
	//Basque
	EU Lang = "eu"
	//Belarusian
	Belarusian Lang = "be"
	//Belarusian
	BE Lang = "be"
	//Bengali
	Bengali Lang = "bn"
	//Bengali
	BN Lang = "bn"
	//Bosnian
	Bosnian Lang = "bs"
	//Bosnian
	BS Lang = "bs"
	//Bulgarian
	Bulgarian Lang = "bg"
	//Bulgarian
	BG Lang = "bg"
	//Catalan
	Catalan Lang = "ca"
	//Catalan
	CA Lang = "ca"
	//Cebuano
	Cebuano Lang = "ceb"
	//Cebuano
	CEB Lang = "ceb"
	//Chinese Simplified
	Chinese_S Lang = "zh-CN"
	//Chinese Simplified
	ZH_CN Lang = "zh-CN"
	//Chinese Traditional
	Chinese_T Lang = "zh-TW"
	//Chinese Traditional
	ZH_TW Lang = "zh-TW"
	//Croatian
	Croatian Lang = "hr"
	//Croatian
	HR Lang = "hr"
	//Czech
	Czech Lang = "cs"
	//Czech
	CS Lang = "cs"
	//Danish
	Danish Lang = "da"
	//Danish
	DA Lang = "da"
	//Dutch
	Dutch Lang = "nl"
	//Dutch
	NL Lang = "nl"
	//English
	English Lang = "en"
	//English
	EN Lang = "en"
	//Esperanto
	Esperanto Lang = "eo"
	//Esperanto
	EO Lang = "eo"
	//Estonian
	Estonian Lang = "et"
	//Estonian
	ET Lang = "et"
	//Fillipino
	Fillipino Lang = "tl"
	//Fillipino
	TL Lang = "tl"
	//Finnish
	Finnish Lang = "fi"
	//Finnish
	FI Lang = "fi"
	//French
	French Lang = "fr"
	//French
	FR Lang = "fr"
	//Galician
	Galician Lang = "gl"
	//Galician
	GL Lang = "gl"
	//Georgian
	Georgian Lang = "ka"
	//Georgian
	KA Lang = "ka"
	//German
	German Lang = "de"
	//German
	DE Lang = "de"
	//Greek
	Greek Lang = "el"
	//Greek
	EL Lang = "el"
	//Gujarati
	Gujarati Lang = "gu"
	//Gujarati
	GU Lang = "gu"
	//Haitian Creole
	Haitian Lang = "ht"
	//Haitian Creole
	HT Lang = "ht"
	//Hausa
	Hausa Lang = "ha"
	//Hausa
	HA Lang = "ha"
	//Hebrew
	Hebrew Lang = "iw"
	//Hebrew
	IW Lang = "iw"
	//Hindi
	Hindi Lang = "hi"
	//Hindi
	HI Lang = "hi"
	//Hmong
	Hmong Lang = "hmn"
	//Hmong
	HMN Lang = "hmn"
	//Hungarian
	Hungarian Lang = "hu"
	//Hungarian
	HU Lang = "hu"
	//Icelandic
	Icelandic Lang = "is"
	//Icelandic
	IS Lang = "is"
	//Igbo
	Igbo Lang = "ig"
	//Igbo
	IG Lang = "ig"
	//Indonesian
	Indonesian Lang = "id"
	//Indonesian
	ID Lang = "id"
	//Irish
	Irish Lang = "ga"
	//Irish
	GA Lang = "ga"
	//Italian
	Italian Lang = "it"
	//Italian
	IT Lang = "it"
	//Japanese
	Japanese Lang = "ja"
	//Japanese
	JA Lang = "ja"
	//Javanese
	Javanese Lang = "jw"
	//Javanese
	JW Lang = "jw"
	//Kannada
	Kannada Lang = "kn"
	//Kannada
	KN Lang = "kn"
	//Khmer
	Khmer Lang = "km"
	//Khmer
	KM Lang = "km"
	//Korean
	Korean Lang = "ko"
	//Korean
	KO Lang = "ko"
	//Lao
	Lao Lang = "lo"
	//Lao
	LO Lang = "lo"
	//Latin
	Latin Lang = "la"
	//Latin
	LA Lang = "la"
	//Latvian
	Latvian Lang = "lv"
	//Latvian
	LV Lang = "lv"
	//Lithuanian
	Lithuanian Lang = "lt"
	//Lithuanian
	LT Lang = "lt"
	//Macedonian
	Macedonian Lang = "mk"
	//Macedonian
	MK Lang = "mk"
	//Malay
	Malay Lang = "ms"
	//Malay
	MS Lang = "ms"
	//Maltese
	Maltese Lang = "mt"
	//Maltese
	MT Lang = "mt"
	//Maori
	Maori Lang = "mi"
	//Maori
	MI Lang = "mi"
	//Marathi
	Marathi Lang = "mr"
	//Marathi
	MR Lang = "mr"
	//Mongolian
	Mongolian Lang = "mn"
	//Mongolian
	MN Lang = "mn"
	//Nepali
	Nepali Lang = "ne"
	//Nepali
	NE Lang = "ne"
	//Norwegian
	Norwegian Lang = "no"
	//Norwegian
	NO Lang = "no"
	//Persian
	Persian Lang = "fa"
	//Persian
	FA Lang = "fa"
	//Polish
	Polish Lang = "pl"
	//Polish
	PL Lang = "pl"
	//Portuguese
	Portuguese Lang = "pt"
	//Portuguese
	PT Lang = "pt"
	//Punjabi
	Punjabi Lang = "pa"
	//Punjabi
	PA Lang = "pa"
	//Romanian
	Romanian Lang = "ro"
	//Romanian
	RO Lang = "ro"
	//Russian
	Russian Lang = "ru"
	//Russian
	RU Lang = "ru"
	//Serbian
	Serbian Lang = "sr"
	//Serbian
	SR Lang = "sr"
	//Slovak
	Slovak Lang = "sk"
	//Slovak
	SK Lang = "sk"
	//Slovenian
	Slovenian Lang = "sl"
	//Slovenian
	SL Lang = "sl"
	//Somali
	Somali Lang = "so"
	//Somali
	SO Lang = "so"
	//Spanish
	Spanish Lang = "es"
	//Spanish
	ES Lang = "es"
	//Swahili
	Swahili Lang = "sw"
	//Swahili
	SW Lang = "sw"
	//Swedish
	Swedish Lang = "sv"
	//Swedish
	SV Lang = "sv"
	//Tamil
	Tamil Lang = "ta"
	//Tamil
	TA Lang = "ta"
	//Telugu
	Telugu Lang = "te"
	//Telugu
	TE Lang = "te"
	//Thai
	Thai Lang = "th"
	//Thai
	TH Lang = "th"
	//Turkish
	Turkish Lang = "tr"
	//Turkish
	TR Lang = "tr"
	//Ukranian
	Ukranian Lang = "uk"
	//Ukranian
	UK Lang = "uk"
	//Urdu
	Urdu Lang = "ur"
	//Urdu
	UR Lang = "ur"
	//Vietnamese
	Vietnamese Lang = "vi"
	//Vietnamese
	VI Lang = "vi"
	//Weish
	Weish Lang = "cy"
	//Weish
	CY Lang = "cy"
	//Yiddish
	Yiddish Lang = "yi"
	//Yiddish
	YI Lang = "yi"
	//Yoruba
	Yoruba Lang = "yo"
	//Yoruba
	YO Lang = "yo"
	//Zulu
	Zulu Lang = "zu"
	//Zulu
	ZU Lang = "zu"

	//Auto identify language
	AUTO      Lang = "auto"
	AUTOMATIC Lang = "auto"
)

//=======================================================================
//							Types
//=======================================================================

type Lang string
