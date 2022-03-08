package translator

import (
	"fmt"
	"log"
)

type Language string

const (
	English Language = "en"
	German  Language = "de"
)

var lang Language

func SetLanguage(language Language) {
	if language != English && language != German {
		panic(fmt.Errorf("unsupported language '%s'", string(language)))
	}
	lang = language
}

// G gets the translation for a string depending on the
// setting of Translator.Language.
//
// Panics, if stringID does not exists and falls back to
// english, if the other locale translation does not exist
//
// Allows the autofilling of string formatters with fmt.Sprintf.
func G(stringID string, a ...interface{}) string {
	translationSet, ok := translations[stringID]
	if !ok {
		panic(fmt.Errorf("translation set for stringID '%s' not found", stringID))
	}
	translation, ok := translationSet[lang]
	if !ok {
		// try 'en'
		translation, ok = translationSet[English]
		if !ok {
			panic(fmt.Errorf("no translation found for 'en' after falling back from '%s'", string(lang)))
		}
		log.Printf("Warning: falling back to 'en', as translation '%s' does not exist for '%s'\n", string(lang), stringID)
	}
	return fmt.Sprintf(translation, a...)
}
