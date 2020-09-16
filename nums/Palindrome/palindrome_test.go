package Palindrome

import (
	"log"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if IsPalindrome(121) {
		log.Println("通过")
	}
	if IsPalindrome(-121) == false {
		log.Println("不通过")
	}
}
