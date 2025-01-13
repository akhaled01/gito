package dns

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type DNSHeader struct {
	ID      uint16 // Transaction ID
	Flags   uint16 // Flags (QR, OpCode, AA, TC, RD, RA, Z, RCODE)
	QDCOUNT uint16 // Number of questions in the message
	ANCOUNT uint16 // Number of answers in the message
	NSCOUNT uint16 // Number of name server records in the message
	ARCOUNT uint16 // Number of additional records in the message
}

type DNSHeaderFlags map[string]uint16

func DecodeHeader(data []byte) (*DNSHeader, error) {
	if len(data) < 12 {
		return nil, ErrPacketTooShort
	}

	header := &DNSHeader{}

	buf := bytes.NewReader(data)

	err := binary.Read(buf, binary.BigEndian, header)
	if err != nil {
		return nil, fmt.Errorf("error reading DNS header: %v", err)
	}

	return header, nil
}

func (h *DNSHeader) GetFlags() DNSHeaderFlags {
	return DNSHeaderFlags{
		"QR":     (h.Flags >> 15) & 0x01,
		"OPCODE": (h.Flags >> 11) & 0x0F,
		"AA":     (h.Flags >> 10) & 0x01,
		"TC":     (h.Flags >> 9) & 0x01,
		"RD":     (h.Flags >> 8) & 0x01,
		"RA":     (h.Flags >> 7) & 0x01,
		"RCODE":  h.Flags & 0x0F,
	}
}
