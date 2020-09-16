package StoI

import (
	"log"
	"testing"
)

func TestMyAtoiOtheri(t *testing.T) {
	num := MyAtoiOther("words and 987")
	if num != 987 {
		log.Println("测试不通过")
		log.Println(num)
		return
	}
	log.Println("测试通过")
	log.Printf("num: %d, %T",num, num)
}
