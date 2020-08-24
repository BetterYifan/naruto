package _3_factory_method

import "testing"

func complete(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var (
		factory OperatorFactory
	)
	factory = PlusOperatorFactory{}
	if complete(factory, 1, 3) != 4 {
		t.Fatal("error")
	}
}
