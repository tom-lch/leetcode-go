package Integer2Roman

import (
	"log"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	if IntToRoman1(3) == "III" {
		log.Println("测试通过")
	}
}
