package util

// python like mod, -1 % 5 == 4, in go -1 % 5 == -1
func Mod(a, b int) int {
	return (a%b + b) % b
}
