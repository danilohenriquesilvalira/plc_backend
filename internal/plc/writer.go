package plc

import (
	"encoding/binary"
	"fmt"
	"math"
)

func (c *Client) WriteTag(dbNumber, byteOffset int, dataType string, value interface{}) error {
	var buf []byte

	switch dataType {
	case "real":
		buf = make([]byte, 4)
		val, ok := value.(float32)
		if !ok {
			return fmt.Errorf("valor deve ser float32")
		}
		binary.BigEndian.PutUint32(buf, math.Float32bits(val))

	case "int":
		buf = make([]byte, 2)
		val, ok := value.(int16)
		if !ok {
			return fmt.Errorf("valor deve ser int16")
		}
		binary.BigEndian.PutUint16(buf, uint16(val))

	case "word":
		buf = make([]byte, 2)
		val, ok := value.(uint16)
		if !ok {
			return fmt.Errorf("valor deve ser uint16")
		}
		binary.BigEndian.PutUint16(buf, val)

	case "bool":
		buf = make([]byte, 1)
		val, ok := value.(bool)
		if !ok {
			return fmt.Errorf("valor deve ser bool")
		}
		if val {
			buf[0] = 0x01
		}

	case "string":
		str, ok := value.(string)
		if !ok {
			return fmt.Errorf("valor deve ser string")
		}
		if len(str) > 254 {
			str = str[:254]
		}
		buf = make([]byte, len(str)+2)
		buf[0] = 254
		buf[1] = byte(len(str))
		copy(buf[2:], str)

	default:
		return fmt.Errorf("tipo não suportado: %s", dataType)
	}

	return c.client.AGWriteDB(dbNumber, byteOffset, len(buf), buf)
}
