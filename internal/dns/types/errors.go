package dns

import "errors"

var (
	ErrPacketTooShort   = errors.New("DNS packet too short")
	ErrInvalidFlagValue = errors.New("value for flag is invalid")
)
