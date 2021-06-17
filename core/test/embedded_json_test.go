package testing

import (
	"fmt"
	"testing"
)

var jsonStr  = `{"basic_info":{"name":"小明","age":12},"job_info":{"skills":["Java","Go","C"]}}`
func TestEmployee_MarshalEasyJSON(t *testing.T) {
	e := EEEdfe{}
	e.UnmarshalJSON([]byte(jsonStr))
	fmt.Println(e)
	if v,err:=e.MarshalJSON();err!=nil {
		t.Error(err)
	}else{
		fmt.Println(string(v))
	}
}
