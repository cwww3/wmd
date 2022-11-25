package tools

func Contains[K comparable](slices []K, s K) bool {
	for i := 0; i < len(slices); i++ {
		if slices[i] == s {
			return true
		}
	}
	return false
}
