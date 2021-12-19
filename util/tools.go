package util

// python like mod, -1 % 5 == 4, in go -1 % 5 == -1
func Mod(a, b int) int {
	return (a%b + b) % b
}

func Max(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func Abs(v int) int {
	if v >= 0 {
		return v
	}
	return -v
}
