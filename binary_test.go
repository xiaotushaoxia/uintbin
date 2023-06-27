package uintbin

import (
	"fmt"
	"math"
	"testing"
)

func TestReadBit(t *testing.T) {

	a := uint32(math.MaxUint32 - 4)

	var vs []bool
	for i := 0; i < 32; i++ {
		vs = append(vs, ReadBit(a, i))
	}
	fmt.Println(vs)

	var k22k uint32
	fmt.Println(Uint2BinaryString(a))
	BinaryString2Uint(Uint2BinaryString(a), &k22k)
	fmt.Println(k22k == a)

	type kk uint64
	var c = kk(255 - 64)
	var vs2 []bool
	for i := 0; i < 32; i++ {
		vs2 = append(vs2, ReadBit(c, i))
	}
	fmt.Println(vs2)
	fmt.Println(Uint2BinaryString(c))
}

func TestPutUint(t *testing.T) {
	binaryString := Uint2BinaryString(uint32(65525))
	fmt.Println(binaryString)
	type kk uint64

	var k kk

	BinaryString2Uint(binaryString, &k)
	fmt.Println(k)

	var k2 uint64

	BinaryString2Uint(binaryString, &k2)
	fmt.Println(k2 == uint64(k))
}
