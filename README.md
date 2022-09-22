# csv-conv

Convert structure and csv to each other

- CSV2Struct converts a 2D slice of string to a slice of struct.
  - You can customize the format of a field by overriding its String method
- Struct2CSV converts a slice of struct to a 2D slice of string.

Both of them would panic if the format of the input is not correct.


使用反射实现结构体与csv的互相转换

参数格式错误时，函数会panic

可以通过修改字段的String方法来自定义输出格式

例子见单测文件
