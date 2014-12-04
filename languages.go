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
	//Afrikaans language
	Afrikaans Lang = iota
	//AF language
	AF
	//Albanian language
	Albanian
	//SQ language
	SQ
	//Arabic language
	Arabic
	//AR language
	AR
	//Armenian language
	Armenian
	//HY language
	HY
	//Azerbaijani language
	Azerbaijani
	//AZ language
	AZ
	//Basque language
	Basque
	//EU language
	EU
	//Belarusian language
	Belarusian
	//BE language
	BE
	//Bengali language
	Bengali
	//BN language
	BN
	//Bosnian language
	Bosnian
	//BS language
	BS
	//Bulgarian language
	Bulgarian
	//BG language
	BG
	//Catalan language
	Catalan
	//CA language
	CA
	//Cebuano language
	Cebuano
	//CEB language
	CEB
	//ChineseS language
	ChineseS
	//ZHCN language
	ZHCN
	//ChineseT language
	ChineseT
	//ZHTW language
	ZHTW
	//Croatian language
	Croatian
	//HR language
	HR
	//Czech language
	Czech
	//CS language
	CS
	//Danish language
	Danish
	//DA language
	DA
	//Dutch language
	Dutch
	//NL language
	NL
	//English language
	English
	//EN language
	EN
	//Esperanto language
	Esperanto
	//EO language
	EO
	//Estonian language
	Estonian
	//ET language
	ET
	//Fillipino language
	Fillipino
	//TL language
	TL
	//Finnish language
	Finnish
	//FI language
	FI
	//French language
	French
	//FR language
	FR
	//Galician language
	Galician
	//GL language
	GL
	//Georgian language
	Georgian
	//KA language
	KA
	//German language
	German
	//DE language
	DE
	//Greek language
	Greek
	//EL language
	EL
	//Gujarati language
	Gujarati
	//GU language
	GU
	//Haitian language
	Haitian
	//HT language
	HT
	//Hausa language
	Hausa
	//HA language
	HA
	//Hebrew language
	Hebrew
	//IW language
	IW
	//Hindi language
	Hindi
	//HI language
	HI
	//Hmong language
	Hmong
	//HMN language
	HMN
	//Hungarian language
	Hungarian
	//HU language
	HU
	//Icelandic language
	Icelandic
	//IS language
	IS
	//Igbo language
	Igbo
	//IG language
	IG
	//Indonesian language
	Indonesian
	//ID language
	ID
	//Irish language
	Irish
	//GA language
	GA
	//Italian language
	Italian
	//IT language
	IT
	//Japanese language
	Japanese
	//JA language
	JA
	//Javanese language
	Javanese
	//JW language
	JW
	//Kannada language
	Kannada
	//KN language
	KN
	//Khmer language
	Khmer
	//KM language
	KM
	//Korean language
	Korean
	//KO language
	KO
	//Lao language
	Lao
	//LO language
	LO
	//Latin language
	Latin
	//LA language
	LA
	//Latvian language
	Latvian
	//LV language
	LV
	//Lithuanian language
	Lithuanian
	//LT language
	LT
	//Macedonian language
	Macedonian
	//MK language
	MK
	//Malay language
	Malay
	//MS language
	MS
	//Maltese language
	Maltese
	//MT language
	MT
	//Maori language
	Maori
	//MI language
	MI
	//Marathi language
	Marathi
	//MR language
	MR
	//Mongolian language
	Mongolian
	//MN language
	MN
	//Nepali language
	Nepali
	//NE language
	NE
	//Norwegian language
	Norwegian
	//NO language
	NO
	//Persian language
	Persian
	//FA language
	FA
	//Polish language
	Polish
	//PL language
	PL
	//Portuguese language
	Portuguese
	//PT language
	PT
	//Punjabi language
	Punjabi
	//PA language
	PA
	//Romanian language
	Romanian
	//RO language
	RO
	//Russian language
	Russian
	//RU language
	RU
	//Serbian language
	Serbian
	//SR language
	SR
	//Slovak language
	Slovak
	//SK language
	SK
	//Slovenian language
	Slovenian
	//SL language
	SL
	//Somali language
	Somali
	//SO language
	SO
	//Spanish language
	Spanish
	//ES language
	ES
	//Swahili language
	Swahili
	//SW language
	SW
	//Swedish language
	Swedish
	//SV language
	SV
	//Tamil language
	Tamil
	//TA language
	TA
	//Telugu language
	Telugu
	//TE language
	TE
	//Thai language
	Thai
	//TH language
	TH
	//Turkish language
	Turkish
	//TR language
	TR
	//Ukranian language
	Ukranian
	//UK language
	UK
	//Urdu language
	Urdu
	//UR language
	UR
	//Vietnamese language
	Vietnamese
	//VI language
	VI
	//Weish language
	Weish
	//CY language
	CY
	//Yiddish language
	Yiddish
	//YI language
	YI
	//Yoruba language
	Yoruba
	//YO language
	YO
	//Zulu language
	Zulu
	//ZU language
	ZU
	//AUTO Automatically recognized lang
	AUTO
	//Automatic recognized Lang
	Automatic
)

//=======================================================================
//						Structs & Types
//=======================================================================

//Lang is a type that represents a ISO 693-1 language in a easy way
type Lang uint8

var langs = [...]string{
	"af",
	"af",
	"sq",
	"sq",
	"ar",
	"ar",
	"hy",
	"hy",
	"az",
	"az",
	"eu",
	"eu",
	"be",
	"be",
	"bn",
	"bn",
	"bs",
	"bs",
	"bg",
	"bg",
	"ca",
	"ca",
	"ceb",
	"ceb",
	"zh-CN",
	"zh-CN",
	"zh-TW",
	"zh-TW",
	"hr",
	"hr",
	"cs",
	"cs",
	"da",
	"da",
	"nl",
	"nl",
	"en",
	"en",
	"eo",
	"eo",
	"et",
	"et",
	"tl",
	"tl",
	"fi",
	"fi",
	"fr",
	"fr",
	"gl",
	"gl",
	"ka",
	"ka",
	"de",
	"de",
	"el",
	"el",
	"gu",
	"gu",
	"ht",
	"ht",
	"ha",
	"ha",
	"iw",
	"iw",
	"hi",
	"hi",
	"hmn",
	"hmn",
	"hu",
	"hu",
	"is",
	"is",
	"ig",
	"ig",
	"id",
	"id",
	"ga",
	"ga",
	"it",
	"it",
	"ja",
	"ja",
	"jw",
	"jw",
	"kn",
	"kn",
	"km",
	"km",
	"ko",
	"ko",
	"lo",
	"lo",
	"la",
	"la",
	"lv",
	"lv",
	"lt",
	"lt",
	"mk",
	"mk",
	"ms",
	"ms",
	"mt",
	"mt",
	"mi",
	"mi",
	"mr",
	"mr",
	"mn",
	"mn",
	"ne",
	"ne",
	"no",
	"no",
	"fa",
	"fa",
	"pl",
	"pl",
	"pt",
	"pt",
	"pa",
	"pa",
	"ro",
	"ro",
	"ru",
	"ru",
	"sr",
	"sr",
	"sk",
	"sk",
	"sl",
	"sl",
	"so",
	"so",
	"es",
	"es",
	"sw",
	"sw",
	"sv",
	"sv",
	"ta",
	"ta",
	"te",
	"te",
	"th",
	"th",
	"tr",
	"tr",
	"uk",
	"uk",
	"ur",
	"ur",
	"vi",
	"vi",
	"cy",
	"cy",
	"yi",
	"yi",
	"yo",
	"yo",
	"zu",
	"zu",
	"auto",
	"auto"}

//=======================================================================
//							Funcs
//=======================================================================

func (l Lang) valid() bool {
	return uint8(l) < uint8(len(langs))
}

//Stringer interface for Lang type
func (l Lang) String() string {
	if !l.valid() {
		return "invalid!"
	}
	return langs[l]
}

//IsLang is a function to check if the provided string is a valid Lang
func IsLang(lang string) bool {
	for _, val := range langs {
		if lang == val {
			return true
		}
	}
	return false
}

type private interface {
	private()
}

func (l Lang) private() {}
