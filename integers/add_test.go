package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expect := 4

	if sum != expect {
		t.Errorf("expect '%d' but get '%d'", expect, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum) //如果没有Output，那么该段代码虽然会被编译，但是不会被执行
	// Output: 6
}
