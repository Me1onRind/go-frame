package slice

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type Uint32Slice []uint32

func (u *Uint32Slice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal []Uint32 value:", value))
	}

	*u = SplitUint32Slice(string(bytes), ",")
	return nil
}

func (u Uint32Slice) Value() (driver.Value, error) {
	if len(u) == 0 {
		return "", nil
	}
	return JoinUint32Slice(u, ","), nil
}
