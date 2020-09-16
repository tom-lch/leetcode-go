package ZigZag_Conversion

import (
	"log"
	"testing"
)

func TestConvert(t *testing.T) {
	s := "LEETCODEISHIRING"
	numRows := 3
	if Convert(s, numRows) == "LCIRETOESIIGEDHN" {
		log.Println("测试1通过测试")
	} else {
		log.Println("测试1 失败")
	}
	numRows = 4
	if Convert(s, numRows) == "LDREOEIIECIHNTSG" {
		log.Println("测试2通过测试")
	} else {
		log.Println("测试2 失败")
	}

}
