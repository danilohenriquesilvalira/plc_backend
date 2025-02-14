package plc

import (
	"encoding/binary"
	"fmt"
	"math"
)

func (c *Client) ReadTag(dbNumber, byteOffset int, dataType string) (interface{}, error) {
	var size int

	switch dataType {
	case "real":
		size = 4
	case "int":
		size = 2
	case "word":
		size = 2
	case "bool":
		size = 1
	case "string":
		size = 256
	default:
		return nil, fmt.Errorf("tipo não suportado: %s", dataType)
	}

	buf := make([]byte, size)
	if err := c.client.AGReadDB(dbNumber, byteOffset, size, buf); err != nil {
		return nil, err
	}

	switch dataType {
	case "real":
		return math.Float32frombits(binary.BigEndian.Uint32(buf)), nil
	case "int":
		return int16(binary.BigEndian.Uint16(buf)), nil
	case "word":
		return binary.BigEndian.Uint16(buf), nil
	case "bool":
		return (buf[0] & 0x01) != 0, nil
	case "string":
		strLen := int(buf[1])
		if strLen > 254 {
			strLen = 254
		}
		return string(buf[2 : 2+strLen]), nil
	}
	return nil, fmt.Errorf("tipo não suportado: %s", dataType)
}
