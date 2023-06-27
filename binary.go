package uintbin

import (
	"fmt"
)

func Uint2BinaryString(u any) string {
	switch v := toInt(u).(type) {
	case uint64:
		return toBinaryString(v, 64)
	case uint32:
		return toBinaryString(uint64(v), 32)
	case uint16:
		return toBinaryString(uint64(v), 16)
	case uint8:
		return toBinaryString(uint64(v), 8)
	case uint:
		return toBinaryString(uint64(v), intSize)
	case int64:
		return toBinaryString(uint64(v), 64)
	case int32:
		return toBinaryString(uint64(v), 32)
	case int16:
		return toBinaryString(uint64(v), 16)
	case int8:
		return toBinaryString(uint64(v), 8)
	case int:
		return toBinaryString(uint64(v), intSize)
	default:
		panic(fmt.Sprintf("not support type %T", u))
	}
}

// BinaryString2Uint 因为result是输入的 所以不需要支持int类型了
func BinaryString2Uint(s string, result any) {
	set2uintPtr(sting2uint64(s), result)
}

func ReadBit[T ~uint16 | ~uint8 | ~uint32 | ~uint64 | ~uint](u T, n int) bool {
	return (u>>n)&1 != 0
}

func ReadBitAny(u any, n int) bool {
	switch uu := toInt(u).(type) {
	case uint64:
		return ReadBit(uu, n)
	case uint32:
		return ReadBit(uu, n)
	case uint16:
		return ReadBit(uu, n)
	case uint8:
		return ReadBit(uu, n)
	case uint:
		return ReadBit(uu, n)
	case int64:
		return ReadBit(uint64(uu), n)
	case int32:
		return ReadBit(uint32(uu), n)
	case int16:
		return ReadBit(uint16(uu), n)
	case int8:
		return ReadBit(uint8(uu), n)
	case int:
		return ReadBit(uint(uu), n)
	default:
		panic(fmt.Sprintf("not support type %T", u))
	}
	// 为了处理类型别名的情况，类似泛型里面的~uint32
	// type kk uint32
	// v := kk(33)
	// ReadBitAny(v, 3)
}
