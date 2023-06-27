package uintbin

import (
	"fmt"
	"testing"
)

type kk2 uint32

func TestUint2Bytes(t *testing.T) {
	a := uint16(65531)

	bytes := Uint2Bytes(a)
	fmt.Println(bytes)

	fmt.Println(bytes2uint64(bytes))

	b := kk2(65531*19 + 2)
	c := kk2(1)
	bytes = Uint2Bytes(b)
	fmt.Println(bytes)

	fmt.Println(bytes2uint64(bytes))
	Bytes2Uint(bytes, &c)
	fmt.Println(c)
}
