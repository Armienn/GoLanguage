package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Armienn/GoLanguage/language"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lang := language.GetDansk()
	word := lang.RandomWord(0)
	fmt.Println(lang.GetRepresentation(word))
}
