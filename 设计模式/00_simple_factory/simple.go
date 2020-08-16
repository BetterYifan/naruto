package simplefactory

import "fmt"

/**
Go没有构造函数的说法，所以一般通过NewXXX来初始化相关类(struct)。NewXXX函数返回接口时就是简单工厂。
*/

type Animal interface {
	Bite(name string) string
}

func NewAnimal(t int) Animal {
	if t == 1 {
		return &Dog{}
	} else if t == 2 {
		return &Cat{}
	}
	return nil
}

type Dog struct{}

func (d *Dog) Bite(name string) string {
	return fmt.Sprintf("Dog:%s", name)
}

type Cat struct{}

func (c *Cat) Bite(name string) string {
	return fmt.Sprintf("Cat: %s", name)
}
