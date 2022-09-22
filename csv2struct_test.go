package csv_conv

import (
	"reflect"
	"testing"
)

type test struct {
	ID        int     `json:"id" csv:"ID" csv_order:"1"`
	Name      string  `json:"name" csv:"名字" csv_order:"2"`
	Weight    float64 `json:"weight" csv:"体重" csv_order:"4"`
	UseGolang bool    `json:"use_golang" csv:"是否使用Golang" csv_order:"3"`
}

var tests = []test{
	{
		ID:        1,
		Name:      "test1",
		Weight:    1.1,
		UseGolang: true,
	},
	{
		ID:        2,
		Name:      "test2",
		Weight:    2.2,
		UseGolang: false,
	},
}

func TestCSV2Struct(t *testing.T) {
	records := Struct2CSV(tests)
	tp := reflect.TypeOf(test{})
	tmp := CSV2Struct(records, tp)
	tests := tmp.([]test)
	for _, v := range tests {
		t.Logf("%#v\n", v)
	}
}
