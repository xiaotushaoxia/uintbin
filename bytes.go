package uintbin

import (
	"encoding/binary"
	"fmt"
)

func Uint2Bytes(u any) []byte {
	var result []byte
	switch v := toInt(u).(type) {
	case uint64:
		result = make([]byte, 8)
		binary.BigEndian.PutUint64(result, v)
	case uint32:
		result = make([]byte, 4)
		binary.BigEndian.PutUint32(result, v)
	case uint16:
		result = make([]byte, 2)
		binary.BigEndian.PutUint16(result, v)
	case uint8:
		result = make([]byte, 1)
		result[0] = v
	case uint:
		if intSize == 32 {
			result = make([]byte, 4)
			binary.BigEndian.PutUint32(result, uint32(v))
		} else {
			result = make([]byte, 8)
			binary.BigEndian.PutUint64(result, uint64(v))
		}
	case int64:
		result = make([]byte, 8)
		binary.BigEndian.PutUint64(result, uint64(v))
	case int32:
		result = make([]byte, 4)
		binary.BigEndian.PutUint32(result, uint32(v))
	case int16:
		result = make([]byte, 2)
		binary.BigEndian.PutUint16(result, uint16(v))
	case int8:
		result = make([]byte, 1)
		result[0] = uint8(v)
	case int:
		if intSize == 32 {
			result = make([]byte, 4)
			binary.BigEndian.PutUint32(result, uint32(v))
		} else {
			result = make([]byte, 8)
			binary.BigEndian.PutUint64(result, uint64(v))
		}
	default:
		panic(fmt.Sprintf("not support type %T", u))
	}
	return result
}

// Bytes2Uint 因为result是输入的 所以不需要支持int类型了
func Bytes2Uint(s []byte, result any) {
	set2uintPtr(bytes2uint64(s), result)
}
