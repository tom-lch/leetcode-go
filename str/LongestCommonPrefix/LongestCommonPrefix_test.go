package LongestCommonPrefix

import (
	"log"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	s := []string{"dfgdfg", "sfdgd", "sf", "regerg", ""}
	if LongestCommonPrefix(s) == "sd" {
		log.Println("tongg")
	}
}
