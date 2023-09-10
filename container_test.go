package widevine

import (
	"fmt"
	"testing"
)

var test_container = Container{
	ID: []byte{
		0xbd, 0xfa, 0x4d, 0x6c, 0xdb, 0x39, 0x70, 0x2e,
		0x5b, 0x68, 0x1f, 0x90, 0x61, 0x7f, 0x9a, 0x7e,
	},
	Key: []byte{
		0xe2, 0x58, 0xb6, 0x7d, 0x75, 0x42, 0x0, 0x66,
		0xc8, 0x42, 0x4b, 0xd1, 0x42, 0xf8, 0x45, 0x65,
	},
}

func Test_Container(t *testing.T) {
	fmt.Println(test_container)
}
