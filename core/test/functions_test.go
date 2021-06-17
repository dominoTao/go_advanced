package testing

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1,2,3}
	expected := [...]int{1,4,9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret!=expected[i] {
			t.Errorf("input is %d, the expected is %d, the actual %d", inputs[i], expected[i], ret)
		}

	}
}

func TestErrorInCode(t *testing.T) {  // 继续执行后续代码
	fmt.Println("start")
	t.Error("Error")
	fmt.Println("end")
}
func TestFailInCode(t *testing.T) {  // 中断后续代码
	fmt.Println("start")
	t.Fatal("Error")
	fmt.Println("end")
}

func TestSquareAssert(t *testing.T) {
	inputs := [...]int{1,2,3}
	expected := [...]int{1,4,9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
		//if ret!=expected[i] {
		//	t.Errorf("input is %d, the expected is %d, the actual %d", inputs[i], expected[i], ret)
		//}

	}
}