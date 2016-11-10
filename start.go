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

	for _, info := range lang.Sounds {
		fmt.Print(dansk.GetRepresentation(info.Sound))
	}
	fmt.Println()
	for _, info := range lang.Sounds {
		fmt.Print(ipa.GetRepresentation(info.Sound))
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		word := lang.RandomWord(0)
		fmt.Println(lang.GetWordRepresentation(word))
		fmt.Println(dansk.GetWordRepresentation(word))
		fmt.Println(ipa.GetWordRepresentation(word))
	}
}
