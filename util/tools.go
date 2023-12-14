package util

import "math"

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

// Lcm returns the lcm of two numbers using the fact that lcm(a,b) * gcd(a,b) = | a * b |
func Lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

// Iterative Faster iterative version of GcdRecursive without holding up too much of the stack
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
