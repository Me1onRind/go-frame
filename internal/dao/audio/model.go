package audio

import (
	"go-frame/internal/lib/base"
)

// StoreStatus
const (
	Init uint8 = iota
	Sync
	UnSync
)

type Audio struct {
	*base.Model

	Filename   string
	StoreIndex string
	Status     uint8
}
