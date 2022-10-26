package main

import "testing"

// 编写测试
// 1. xxx_test.go
// 2. 测试函数的命名必须以 Test 开始
// 3. 测试函数只接受一个参数 t *testing.T
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
