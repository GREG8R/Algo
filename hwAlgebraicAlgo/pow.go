package main

func PowOfComposition(number float64, pow int) float64 {
	if number == 0 {
		return float64(0)
	}
	if pow == 0 {
		return float64(1)
	}
	result := number
	for i := 1; i < pow; i++ {
		result = result * number
	}
	return result
}

func PowOfPowTwo(number float64, pow int) float64 {
	if number == 0 {
		return float64(0)
	}
	if pow == 0 {
		return float64(1)
	}
	result := number
	k := 1
	for i := 1; i < pow/2; i *= 2 {
		result *= result
		k *= 2
	}

	for i := k; i < pow; i++ {
		result *= number
	}

	return result
}

func FastPow(number float64, pow int) float64 {
	if number == 0 {
		return float64(0)
	}
	if pow == 0 {
		return float64(1)
	}

	result := float64(1)
	for pow > 1 {
		if pow%2 == 1 {
			result *= number
		}
		number *= number
		pow /= 2
	}

	if pow > 0 {
		result *= number
	}
	return result
}
