package uintbin

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

var (
	tUint64 = reflect.TypeOf(uint64(1))
	tUint32 = reflect.TypeOf(uint32(1))
	tUint16 = reflect.TypeOf(uint16(1))
	tUint8  = reflect.TypeOf(uint8(1))
	tUint   = reflect.TypeOf(uint(1))
	tInt64  = reflect.TypeOf(int64(1))
	tInt32  = reflect.TypeOf(int32(1))
	tInt16  = reflect.TypeOf(int16(1))
	tInt8   = reflect.TypeOf(int8(1))
	tInt    = reflect.TypeOf(int(1))
)

const intSize = 32 << (^uint(0) >> 63)

// toInt 用于转换衍生数据
// 为了处理类型别名的情况，类似泛型里面的~uint32
// type kk uint32
// v := kk(33)
// toInt(v) // uint32(33)
func toInt(u any) any {
	var targetType reflect.Type
	switch reflect.TypeOf(u).Kind() {
	case reflect.Uint64:
		targetType = tUint64 // 获取他的原始类型
	case reflect.Uint32:
		targetType = tUint32
	case reflect.Uint16:
		targetType = tUint16
	case reflect.Uint8:
		targetType = tUint8
	case reflect.Uint:
		targetType = tUint
	case reflect.Int64:
		targetType = tInt64
	case reflect.Int32:
		targetType = tInt32
	case reflect.Int16:
		targetType = tInt16
	case reflect.Int8:
		targetType = tInt8
	case reflect.Int:
		targetType = tInt
	default:
		panic(fmt.Sprintf("%T is not int or uint, can't convert to uint)", u))
	}
	return reflect.ValueOf(u).Convert(targetType).Interface()
}

func toBinaryString(v uint64, n int) string {
	vv := fmt.Sprintf("%b", v)
	l := len(vv)
	if l == n {
		return vv
	}
	if l > n {
		return vv[l-n:]
	}
	var result = make([]byte, n-l)
	for i := 0; i < n-l; i++ {
		result[i] = '0'
	}
	return string(result) + vv
	// 换了一下可以直接用%b
	//var result = make([]byte, n)
	//for i := 0; i < n; i++ {
	//	if v&1 == 1 {
	//		result[n-i-1] = '1'
	//	} else {
	//		result[n-i-1] = '0'
	//	}
	//	v >>= 1
	//}
	//return string(result)
}

func sting2uint64(s string) uint64 {
	if len(s) > 64 {
		s = s[len(s)-64:]
	}
	var v uint64
	var k uint64 = 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			v += k
		}
		k *= 2
	}
	return v
}

func bytes2uint64(b []byte) uint64 {
	var tmp uint64
	if len(b) < 8 {
		needFill := make([]byte, 8-len(b))
		b = append(needFill, b...)
	}
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &tmp)
	if err != nil {
		panic(fmt.Errorf("bytes to uint64 fail %s", err))
	}
	return tmp
}

func set2uintPtr(v uint64, result any) {
	switch data := result.(type) {
	case *uint8:
		*data = uint8(v)
	case *uint16:
		*data = uint16(v)
	case *uint32:
		*data = uint32(v)
	case *uint64:
		*data = v
	case *uint:
		*data = uint(v)
	default:
		// 处理衍生类型 类似~int32
		value := reflect.ValueOf(result)
		if value.Kind() != reflect.Pointer {
			panic(fmt.Sprintf("BinaryString2Uint arg must be pointer, get %T", result))
		}
		elemType := reflect.TypeOf(result).Elem()

		var vv reflect.Value
		switch elemType.Kind() {
		case reflect.Uint64:
			vv = reflect.ValueOf(v)
		case reflect.Uint32:
			vv = reflect.ValueOf(uint32(v))
		case reflect.Uint16:
			vv = reflect.ValueOf(uint16(v))
		case reflect.Uint8:
			vv = reflect.ValueOf(uint8(v))
		case reflect.Uint:
			vv = reflect.ValueOf(uint(v))
		default:
			panic(fmt.Sprintf("not support type %T", result))
		}
		value.Elem().Set(vv.Convert(elemType))
		//unsafePtr := unsafe.Pointer(value.Pointer())
		//reflect.NewAt(elemType, unsafePtr).Elem().Set(vv.Convert(elemType))
	}

}
