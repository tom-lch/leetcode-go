package PermutationsII

import (
	"log"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	nums := []int{1, 1, 0}
	arrss := PermuteUnique(nums)
	log.Println(arrss)
}
