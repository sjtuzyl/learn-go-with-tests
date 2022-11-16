package ref

import "reflect"

// 反射: 计算中的反射提供了程序检查自身结构体的能力
// 除非真的需要, 否则不要使用反射

// switch和select的异同
// 1.每个switch后必须跟随一个条件判断，而select没有
// 2.switch中的case为与枚举值比较，select中的case必须是一个对channel的读或者写操作
// 3.switch分支顺序执行，select同时多分支满足则随机选取

func Walk(x interface{}, fn func(input string)) {
	// fn("KRDOG")
	val := getValue(x)

	// 1.虽然可行，但是恶心
	// if val.Kind() == reflect.Slice {
	// 	for i := 0; i < val.Len(); i++ {
	// 		Walk(val.Index(i).Interface(), fn)
	// 	}
	// 	return
	// }

	// for i := 0; i < val.NumField(); i++ {
	// 	field := val.Field(i)

	// 	// if field.Kind() == reflect.String {
	// 	// 	fn(field.String())
	// 	// } else if field.Kind() == reflect.Struct {
	// 	// 	Walk(field.Interface(), fn)
	// 	// }
	// 	switch field.Kind() {
	// 	case reflect.String:
	// 		fn(field.String())
	// 	case reflect.Struct:
	// 		Walk(field.Interface(), fn)
	// 	}
	// }

	// 2.还可以做得更好
	// switch val.Kind() {
	// case reflect.Struct:
	// 	for i := 0; i < val.NumField(); i++ {
	// 		Walk(val.Field(i).Interface(), fn)
	// 	}
	// case reflect.Slice:
	// 	for i := 0; i < val.Len(); i++ {
	// 		Walk(val.Index(i).Interface(), fn) //
	// 	}
	// case reflect.String:
	// 	fn(val.String())
	// }

	// 3.再次重构
	// numOfValues := 0
	// var getField func(int) reflect.Value
	// switch val.Kind() {
	// case reflect.Struct:
	// 	numOfValues = val.NumField()
	// 	getField = val.Field
	// case reflect.Slice, reflect.Array:
	// 	numOfValues = val.Len()
	// 	getField = val.Index
	// case reflect.String:
	// 	fn(val.String())
	// case reflect.Map:
	// 	for _, key := range val.MapKeys() {
	// 		Walk(val.MapIndex(key).Interface(), fn)
	// 	}
	// }

	// for i := 0; i < numOfValues; i++ {
	// 	Walk(getField(i).Interface(), fn)
	// }

	waleValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			waleValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			waleValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			waleValue(val.MapIndex(key))
		}
	case reflect.String:
		fn(val.String())
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
