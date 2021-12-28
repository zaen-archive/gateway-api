package utils

func ArrayContain(arr []string, contain string) bool {
	for _, s := range arr {
		if s == contain {
			return true
		}
	}

	return false
}

// EmptyStrig
func EmptyString(s *string) bool {
	return len(*s) == 0
}
