package dns

import (
	"bytes"
	"encoding/binary"
)

type DNSAnswer struct {
	NAME      string
	TYPE      uint16
	CLASS     uint16
	TTL       uint16
	RDLENGTH  uint16
	RDATA     uint16
	PREFRENCE uint16
	EXCHANCE  uint16
}

func (a *DNSAnswer) Encode() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, a); err != nil {
		return nil, err
	}

	return buf, nil
}
