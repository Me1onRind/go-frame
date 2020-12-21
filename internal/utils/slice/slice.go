package slice

func InSliceUint8(target uint8, slice []uint8) bool {
	for _, v := range slice {
		if target == v {
			return true
		}
	}
	return false
}
