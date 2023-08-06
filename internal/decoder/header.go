package decoder

import (
	"bytes"
	"udp/models"
)

// Декодирует Header
func DecodeHeader(buf *bytes.Buffer, header *models.Header) error {
	return BinaryDecoder(buf, header)
}
