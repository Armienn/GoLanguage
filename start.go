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
	for i := 0; i < 10; i++ {
		word := lang.RandomWord(0)
		fmt.Println(lang.GetRepresentation(word))
		fmt.Println(dansk.GetRepresentation(word))
	}
}
