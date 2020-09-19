package Integer2Roman



func IntToRoman(num int) string {
	housands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return housands[num/1000] + hundreds[num%1000/100] + tens[num%1000%100/10] + ones[num%10]
}

func IntToRoman1(num int) string {
	RR := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	values := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	var s string
	for i := 0; i < 13; i++ {
		for num >= values[i] {
			num -= values[i]
			s += RR[i]
		}
	}
	return s
}