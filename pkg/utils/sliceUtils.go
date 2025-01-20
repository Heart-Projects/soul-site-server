package utils

// Contains 泛型函数，判断切片中是否包含某个函数
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
