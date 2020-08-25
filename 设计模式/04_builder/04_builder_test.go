package _4_builder

import "testing"

func TestBuild1Instance(t *testing.T) {
	builder1 := &Build1Instance{}
	director := NewDirector(builder1)
	director.Create()
	res := builder1.GetResult()
	if res != "ABC" {
		t.Fatal("Error: ", res)
	}
}
