package DivideTwoIntegers

import (
	"log"
	"testing"
)

func TestDivide(t *testing.T) {
	if Divide(10, 3) == 3 {
		log.Println(true)
	}
	log.Println(Divide(10, -3))
	log.Println(Divide(10, -1))
	log.Println(Divide(1<<32, -1))
}
