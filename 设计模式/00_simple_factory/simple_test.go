package simplefactory

import "testing"

//go test -v simple_test.go simple.go

func TestType1(t *testing.T) {
	animal := NewAnimal(1)
	dog := animal.Bite("狗狗")
	if dog != "Dog:狗狗" {
		t.Fatal("test fail")
	}
}
