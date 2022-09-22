package csv_conv

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
Struct2CSV converts a slice of struct to a 2D slice of string
The first row is the header
The following rows are the data
The header is the value of the tag "csv" in the struct field
The order of the columns is the value of the tag "csv_order" in the struct field
The value of the tag "csv_order" must be in range [1, N] where N is the number of fields in the struct
You can customize the format of a field by overriding its String method
Struct2CSV 将一个结构体切片转换为二维字符串切片
第一行是表头 第二行开始是数据
表头是结构体字段的csv标签的值
列的顺序是结构体字段的csv_order标签的值
csv_order标签的值必须在[1, N]范围内，N是结构体的字段数
你可以通过重写字段的String方法来自定义字段的格式
*/
func Struct2CSV(data interface{}) [][]string {
	dataTp := reflect.TypeOf(data)
	if dataTp.Kind() != reflect.Slice {
		panic("data must be a slice")
	}
	if dataTp.Elem().Kind() != reflect.Struct {
		panic("data must be a slice of struct")
	}

	dataVal := reflect.ValueOf(data)
	dataLen := dataVal.Len()
	if dataLen == 0 {
		return nil
	}

	tp := dataVal.Index(0).Type()
	fmt.Println(tp)

	var result [][]string = make([][]string, dataLen+1)

	// 处理表头
	header := make([]string, tp.NumField())
	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		order := getOrderFromTag(field)
		if order < 1 || order > tp.NumField() {
			panic(fmt.Sprintf("csv_order must be in range [1, %d] in field %s", tp.NumField(), field.Name))
		}
		header[order-1] = field.Tag.Get("csv")
	}
	result[0] = header

	for i := 0; i < dataVal.Len(); i++ {
		var row []string = make([]string, tp.NumField())
		// 按照csv_order排序
		for j := 0; j < tp.NumField(); j++ {
			field := tp.Field(j)
			order := getOrderFromTag(field)
			if order < 1 || order > tp.NumField() {
				panic(fmt.Sprintf("csv_order must be in range [1, %d] in field %s", tp.NumField(), field.Name))
			}
			row[order-1] = fmt.Sprintf("%v", dataVal.Index(i).Field(j).Interface())
		}
		result[i+1] = row
	}
	return result
}

func getOrderFromTag(field reflect.StructField) int {
	orderStr, ok := field.Tag.Lookup("csv_order")
	if !ok {
		panic(fmt.Sprintf("csv_order not found in field %s", field.Name))
	}
	order, err := strconv.Atoi(orderStr)
	if err != nil {
		panic(fmt.Sprintf("csv_order must be int in field %s", field.Name))
	}

	return order
}
