package testing

import (
	"reflect"
	"testing"
)

type Employee struct {
	EmployeeID string
	Name string `format:"normal"`
	Age int
}

func (e *Employee)UpdateAge(newVal int)  {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name string
	Age int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	}else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age :",e)
}