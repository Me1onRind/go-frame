package slice

import (
	"bytes"
	"strconv"
	"strings"
)

func InUint8Slice(target uint8, slice []uint8) bool {
	for _, v := range slice {
		if target == v {
			return true
		}
	}
	return false
}

func JoinUint32Slice(slice []uint32, sep string) string {
	buffer := &bytes.Buffer{}
	for k, v := range slice {
		if k > 0 {
			buffer.WriteString(sep)
		}
		buffer.WriteString(strconv.FormatUint(uint64(v), 10))
	}

	return buffer.String()
}

func SplitUint32Slice(src, sep string) []uint32 {
	tmp := strings.Split(src, sep)
	result := make([]uint32, 0, len(tmp))
	for _, v := range tmp {
		number, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			continue
		}
		result = append(result, uint32(number))
	}

	return result
}
