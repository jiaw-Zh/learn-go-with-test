package main

import "fmt"

const englishHelloPrefix = "Hello, "
const chinese = "chinese"
const perchinese = "你好，"
const World = "World"
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

// func Hello(name string, language string) string {
// 	if name == "" {
// 		name = World
// 	}
	// if language == chinese {
	// 	return perchinese + name
	// }
	// if language == spanish {
    //     return spanishHelloPrefix + name
    // }
	// 重构分支语句
// 	switch language {
// 	case chinese:
// 		return perchinese + name
// 	case spanish:
// 		return spanishHelloPrefix + name
// 	}

// 	return englishHelloPrefix + name
// }

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
    switch language {
    case chinese:
        prefix = perchinese
    case spanish:
        prefix = spanishHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return
}



func main() {
	fmt.Println(Hello("charles", ""))
}
