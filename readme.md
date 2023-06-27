# 用于无符号整数和二进制之间的转换 
如255转换成11111111  (uint8转二进制字符串)

65535转换成bytes[255 255]  uint16转字节数组

ReadBit(255,1) -> true

```go
package uintbin

import (
	"fmt"
	"testing"
)

type kk2 uint32

func TestUint2Bytes(t *testing.T) {
	a := uint16(65531)

	bytes := Uint2Bytes(a) 
	fmt.Println(bytes) // [255 251]

	fmt.Println(bytes2uint64(bytes))  // 65531

	b := kk2(65531*19 + 2)
	c := kk2(1)
	bytes = Uint2Bytes(b)  
	fmt.Println(bytes)  // [0 18 255 163]

	fmt.Println(bytes2uint64(bytes)) // 1245091
	Bytes2Uint(bytes, &c)
	fmt.Println(c)  // 1245091
}

func TestReadBit(t *testing.T) {

	a := uint32(math.MaxUint32 - 4)

	var vs []bool
	for i := 0; i < 32; i++ {
		vs = append(vs, ReadBit(a, i))
	}
	fmt.Println(vs)
	//[true true false true true true true true true true true true true true true true true true true true true true true true true true true true true true true true]

	var k22k uint32
	fmt.Println(Uint2BinaryString(a))
	//11111111111111111111111111111011
	
	BinaryString2Uint(Uint2BinaryString(a), &k22k)
	fmt.Println(k22k == a)
	//true

	type kk uint64
	var c = kk(255 - 64)
	var vs2 []bool
	for i := 0; i < 32; i++ {
		vs2 = append(vs2, ReadBit(c, i))
	}
	fmt.Println(vs2)
	//[true true true true true true false true false false false false false false false false false false false false false false false false false false false false false false false false]
	
	fmt.Println(Uint2BinaryString(c))
    //0000000000000000000000000000000000000000000000000000000010111111
}

func TestPutUint(t *testing.T) {
	binaryString := Uint2BinaryString(uint32(65525))
	fmt.Println(binaryString)
	//  00000000000000001111111111110101
	type kk uint64

	var k kk

	BinaryString2Uint(binaryString, &k)
	fmt.Println(k)
    // 65525
	var k2 uint64

	BinaryString2Uint(binaryString, &k2)
	fmt.Println(k2 == uint64(k))
    // true
}
```