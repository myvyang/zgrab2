package ipp_test

import (
	"bytes"
	"fmt"
	"github.com/zmap/zgrab2/modules/ipp"
)

func ExampleAttributeByteString() {
	var buf bytes.Buffer
	if err := ipp.AttributeByteString(0x47, "attributes-charset", "us-ascii", &buf); err == nil {
		fmt.Println(buf.Bytes())
	}
	// Output: [71 0 18 97 116 116 114 105 98 117 116 101 115 45 99 104 97 114 115 101 116 0 8 117 115 45 97 115 99 105 105]
}