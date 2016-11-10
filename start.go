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
	lang := language.RandomLanguage()
	dansk := language.GetDansk()
	ipa := language.GetIpa()
	printAlphabet(dansk, dansk)
	printAlphabet(dansk, ipa)
	printAlphabet(ipa, dansk)
	printAlphabet(ipa, ipa)

	for i := 0; i < 10; i++ {
		lang = language.RandomLanguage()
		printAlphabet(lang, dansk)
		printAlphabet(lang, ipa)
	}

	/*for i := 0; i < 10; i++ {
		word := lang.RandomWord(0)
		fmt.Println(dansk.GetWordRepresentation(word))
		fmt.Println(ipa.GetWordRepresentation(word))
	}*/

}

func printAlphabet(lang *language.Language, representation *language.Language) {
	for _, info := range lang.Sounds {
		fmt.Print(representation.GetRepresentation(info.Sound))
	}
	fmt.Println()
}
