package dns

import (
	"bytes"
	"encoding/binary"
)

type DNSQuestion struct {
	QNAME    string
	Metadata QuestionMetadata
}

type QuestionMetadata struct {
	QTYPE  uint16
	QCLASS uint16
}

func ParseDNSQuestion(data []byte) (*DNSQuestion, error) {
	if len(data) < 12 { // Minimum DNS question size
		return nil, ErrPacketTooShort
	}

	question := &DNSQuestion{}
	offset := 0

	// Parse QNAME
	for {
		if offset >= len(data) {
			return nil, ErrInvalidBytes // Out-of-bounds check
		}

		length := int(data[offset]) // Read length byte
		offset++

		if length == 0 {
			break // End of QNAME
		}

		if offset+length > len(data) {
			return nil, ErrInvalidBytes // Ensure label fits within data
		}

		// Append the label to QNAME
		question.QNAME += string(data[offset : offset+length])
		offset += length // Move past the label

		// Add a dot if not the last label
		if data[offset] != 0x00 {
			question.QNAME += "."
		}
	}

	// Ensure enough bytes remain for metadata (QTYPE + QCLASS)
	if offset+4 > len(data) {
		return nil, ErrInvalidBytes
	}

	// Parse metadata
	raw_metadata := bytes.NewReader(data[offset:])
	if err := binary.Read(raw_metadata, binary.BigEndian, &question.Metadata); err != nil {
		return nil, ErrInvalidBytes
	}

	return question, nil
}
