package translator

import (
	"fmt"
	"log"
)

type Language string

const (
	English Language = "en"
	German  Language = "de"

	fallbackMessage string = "NO TRANSLATION FOUND FOR '%s'"
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
// Returns a fallback message, if stringID does not exist
// and falls back to english, if the other locale translation
// does not exist.
//
// Allows the autofilling of string formatters with fmt.Sprintf.
func G(stringID string, a ...interface{}) string {
	translationSet, ok := translations[stringID]
	if !ok {
		log.Printf("Warning: Translation set for stringID '%s' not found - using fallback message", stringID)
		return fmt.Sprintf(fallbackMessage, stringID)
	}
	translation, ok := translationSet[lang]
	if !ok {
		// try 'en'
		translation, ok = translationSet[English]
		if !ok {
			log.Printf("Warning: No translation found for 'en' after falling back from '%s' - using fallback massage", string(lang))
			return fmt.Sprintf(fallbackMessage, stringID)
		}
		log.Printf("Warning: falling back to 'en', as translation '%s' does not exist for '%s'", string(lang), stringID)
	}
	return fmt.Sprintf(translation, a...)
}
