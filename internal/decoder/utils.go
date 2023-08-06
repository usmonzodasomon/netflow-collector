package decoder

import (
	"encoding/binary"
	"io"
)

// Считывает поля из буфера в бинарном виде
func BinaryDecoder(buf io.Reader, fields ...interface{}) error {
	for _, field := range fields {
		err := binary.Read(buf, binary.BigEndian, field)
		if err != nil {
			return err
		}
	}
	return nil
}
