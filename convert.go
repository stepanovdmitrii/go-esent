package goesent

func boolToUInt(value bool) uint {
	if value {
		return 1
	}
	return 0
}
