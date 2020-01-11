package commons

func IsNilOrEmpty(value interface{}) bool {
	switch value.(type) {
	case nil:
		return true
	case string:
		return value == ""
	case int:
		return value == 0
	}

	return false
}
