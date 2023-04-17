package utils

func Sum(a, b int) (result int) {
	result = a + b
	return
}

func Max(a, b int) (result int) {
	if a < b {
		return b
	}
	return a
}
