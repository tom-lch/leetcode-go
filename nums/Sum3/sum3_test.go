package Sum3

import (
	"log"
	"testing"
)

func TestSum3(t *testing.T) {
	arr := []int{-1,0,1,2,-1,-4}
	ans := Sum3(arr)
	log.Println(ans)
}
