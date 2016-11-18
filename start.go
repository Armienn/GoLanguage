package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Armienn/GoLanguage/grammatics"
	"github.com/Armienn/GoLanguage/language"
	"github.com/Armienn/GoLanguage/phonetics"
)

func main() {
	dansk := language.Dansk{}
	dansk.Phonetics = *phonetics.GetDansk()
	dansk.Words = map[grammatics.Concept]*language.DanishWord{
		"sun":   &language.DanishWord{phonetics.Word{}, "sol", "noun"},
		"shine": &language.DanishWord{phonetics.Word{}, "skinne", "verb"},
	}
	sentences := grammatics.GetSentences()
	for _, sentence := range sentences {
		_, text := dansk.Translate(sentence)
		fmt.Println(text)
	}
}

func printAlphabet(lang *phonetics.Phonetics, representation *phonetics.Phonetics) {
	for _, sound := range lang.Sounds {
		fmt.Print(representation.GetRepresentation(sound))
	}
	fmt.Println()
}

func generateSomePhonetics() {
	rand.Seed(time.Now().UTC().UnixNano())
	//rand.Seed(1)
	lang := phonetics.RandomPhonetics()
	dansk := phonetics.GetDansk()
	ipa := phonetics.GetIpa()
	//printAlphabet(dansk, dansk)
	//printAlphabet(dansk, ipa)
	//printAlphabet(ipa, dansk)
	//printAlphabet(ipa, ipa)

	//for i := 0; i < 10; i++ {
	//lang = language.RandomLanguage()
	lang.Patterns = phonetics.GetMubPatterns()
	printAlphabet(lang, dansk)
	printAlphabet(lang, ipa)
	//}

	for i := 0; i < 20; i++ {
		word := lang.RandomWord(0)
		fmt.Println(dansk.GetWordRepresentation(word))
		fmt.Println(ipa.GetWordRepresentation(word))
	}
}
