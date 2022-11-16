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
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// if field.Kind() == reflect.String {
		// 	fn(field.String())
		// } else if field.Kind() == reflect.Struct {
		// 	Walk(field.Interface(), fn)
		// }
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			Walk(field.Interface(), fn)
		}
	}
}
