package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Armienn/GoLanguage/language"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//rand.Seed(1)
	lang := language.RandomPhonetics()
	dansk := language.GetDansk()
	ipa := language.GetIpa()
	//printAlphabet(dansk, dansk)
	//printAlphabet(dansk, ipa)
	//printAlphabet(ipa, dansk)
	//printAlphabet(ipa, ipa)

	//for i := 0; i < 10; i++ {
	//lang = language.RandomLanguage()
	lang.Patterns = language.GetMubPatterns()
	printAlphabet(lang, dansk)
	printAlphabet(lang, ipa)
	//}

	for i := 0; i < 20; i++ {
		word := lang.RandomWord(0)
		fmt.Println(dansk.GetWordRepresentation(word))
		fmt.Println(ipa.GetWordRepresentation(word))
	}

}

func printAlphabet(lang *language.Phonetics, representation *language.Phonetics) {
	for _, sound := range lang.Sounds {
		fmt.Print(representation.GetRepresentation(sound))
	}
	fmt.Println()
}
