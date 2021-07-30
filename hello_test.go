package main

import (
	"testing"
)

// 初始测试函数
// func TestHello(t *testing.T) {
//     got := Hello("Chris")
//     want := "Hello, Chris"

//     if got != want {
//         t.Errorf("got '%q' want '%q'", got, want)
//     }
// }

// 第一次重构，分支测试语句
// func TestHello(t *testing.T) {

//     t.Run("saying hello to people", func(t *testing.T) {
//         got := Hello("Chris")
//         want := "Hello Chris"

//         if got != want {
//             t.Errorf("got '%q' want '%q'", got, want)
//         }
//     })

//     t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
//         got := Hello("")
//         want := "Hello, World"

//         if got != want {
//             t.Errorf("got '%q' want '%q'", got, want)
//         }
//     })

// }

// 第二次重构，将断言重构为函数，像调用普通函数一样调用断言
func TestHello(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%q' want '%q'", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris","")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'world'", func(t *testing.T) {
        got := Hello("","")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

	t.Run("in chinese", func(t *testing.T) {
		got := Hello("世界", "chinese")
		want := "你好，世界"
		assertCorrectMessage(t, got, want)
	})

}
