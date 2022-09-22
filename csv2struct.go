package csv_conv

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
CSV2Struct converts a 2D slice of string to a slice of struct.
The first row is the header;
The following rows are the data;
The header is the value of the tag "csv" in the struct field;
The order of the columns is the value of the tag "csv_order" in the struct field;
The value of the tag "csv_order" must be in range [1, N] where N is the number of fields in the struct;
Only Integer, Float, Boolean, String are supported.
CSV2Struct 将二维字符串切片转换为结构体切片
第一行是表头 第二行开始是数据
表头是结构体字段的csv标签的值
列的顺序是结构体字段的csv_order标签的值
csv_order标签的值必须在[1, N]范围内，N是结构体的字段数
只支持整数 浮点数 布尔值 字符串
*/
func CSV2Struct(data [][]string, tp reflect.Type) interface{} {
	if len(data) == 0 {
		return nil
	}
	var result = reflect.MakeSlice(reflect.SliceOf(tp), len(data)-1, len(data)-1)
	for i := 1; i < len(data); i++ {
		row := data[i]
		ref := reflect.New(tp)
		for j := 0; j < tp.NumField(); j++ {
			field := tp.Field(j)
			order := getOrderFromTag(field)
			if order < 1 || order > tp.NumField() {
				panic(fmt.Sprintf("csv_order must be in range [1, %d] in field %s", tp.NumField(), field.Name))
			}
			switch field.Type.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				v, err := strconv.Atoi(row[order-1])
				if err != nil {
					panic(err)
				}
				ref.Elem().Field(j).SetInt(int64(v))
			case reflect.String:
				ref.Elem().Field(j).SetString(row[order-1])
			case reflect.Float64, reflect.Float32:
				v, err := strconv.ParseFloat(row[order-1], 64)
				if err != nil {
					panic(err)
				}
				ref.Elem().Field(j).SetFloat(v)
			case reflect.Bool:
				v, err := strconv.ParseBool(row[order-1])
				if err != nil {
					panic(err)
				}
				ref.Elem().Field(j).SetBool(v)
			default:
				panic(fmt.Sprintf("unsupported type %s in field %s", field.Type.Kind(), field.Name))
			}
		}
		result.Index(i - 1).Set(ref.Elem().Convert(tp))
	}
	return result.Interface()
}
