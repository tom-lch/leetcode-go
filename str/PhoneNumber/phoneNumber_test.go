package PhoneNumber

import (
	"log"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	log.Println(LetterCombinations("345"))
	log.Println(LetterCombinations("2345"))
	log.Println(LetterCombinations("23"))
	log.Println(LetterCombinations("9345"))
}