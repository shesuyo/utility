package mathutil

import "math"

func Round(num float64, length int) float64 {
	return round(num, length, true)
}

func UntiRound(num float64, length int) float64 {
	return round(num, length, false)
}

//require x>0 && y>0
func PowInt(x, y int) (result int) {
	if x < 0 || y < 0 {
		panic("powint use range must greater than 0 ")
	}
	result = x
	for i := 0; i < y-1; i++ {
		result *= x
	}
	return
}

func round(num float64, length int, isRound bool) float64 {
	scale := float64(PowInt(10, length))
	num = num * scale
	if isRound {
		num += 0.5
	}
	num = math.Floor(num)
	num /= scale
	return num
}
