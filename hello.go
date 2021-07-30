package main

import "fmt"

const englishHelloPrefix = "Hello, "
const chinese = "chinese"
const perchinese = "你好，"
const World = "World"

func Hello(name string, language string) string {
	if name == "" {
		name = World
	}
	if language == chinese {
		return perchinese + name
	}

	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("charles", ""))
}
