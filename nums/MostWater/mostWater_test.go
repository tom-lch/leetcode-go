package MostWater

import (
	"log"
	"testing"
)

func TestMaxArea(t *testing.T) {
	arr := []int{1,8,6,2,5,4,8,3,7}
	if MaxArea(arr) != 49 {
		log.Println("测试不通过")
	}
	log.Println("测试通过")
}
