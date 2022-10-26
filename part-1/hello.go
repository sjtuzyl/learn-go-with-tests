package main

import "fmt"

// 避免每次使用Hello时创建字符串实例
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
