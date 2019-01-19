package vet

import (
	"testing"
	"fmt"
	"reflect"
)

type Language struct {
	Name string
}

func TestAnalyzeCodes(t *testing.T) {
	instance := new(Language)
	instance.Name = "go"
	dereference_pointer := *instance
	fmt.Println(reflect.TypeOf(dereference_pointer).Kind())
	object := Language{}
	object.Name = "python"
	fmt.Println(object.Name)
}