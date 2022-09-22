package csv_conv

import (
	"strings"
	"testing"
)

func TestStruct2CSV(t *testing.T) {
	records := Struct2CSV(tests)
	for _, v := range records {
		t.Log(strings.Join(v, ","))
	}
}
