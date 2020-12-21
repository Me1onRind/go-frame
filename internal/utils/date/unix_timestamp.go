package date

import (
	"time"
)

func UnixTime() uint32 {
	return uint32(time.Now().Unix())
}
