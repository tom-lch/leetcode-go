package DivideTwoIntegers


func Divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		if dividend > 1<<31-1 {
			return  1<<31-1
		} else if dividend < -1<<31 {
			return -1<<31
		}
		return dividend
	}
	if divisor == -1 {
		if -dividend > 1<<31-1 {
			return  1<<31-1
		} else if -dividend < -1<<31 {
			return -1<<31
		}
		return -dividend
	}
	var flag = 1
	if (dividend > 0 && divisor < 0) || dividend < 0 && divisor > 0 {
		flag = -1
	}
	var c int
	d1, d2 := abs(dividend), abs(divisor)
	for d1-d2 >= 0 {
		d1 -= d2
		c++
	}
	if flag > 0 {
		return c
	}
	return -c

}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}