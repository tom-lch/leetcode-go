package PhoneNumber


var m = map[byte]string{'2':"abc",
	'3':"def",
	'4':"ghi",
	'5':"jkl",
	'6':"mno",
	'7':"pqrs",
	'8':"tuv",
	'9':"wxyz"}

func LetterCombinations(digits string) []string {
	var ans []string
	Combinations(&ans, "", digits, 0)
	return ans
}

func Combinations(arr *[]string, str, dig string, l int) {
	if l == len(dig) {
		*arr = append(*arr, str)
		return
	}
	temp := m[dig[l]]
	for _, v := range temp {
		str += string(v)
		Combinations(arr, str, dig, l+1)
		str = str[:len(str)-1]
	}
}
