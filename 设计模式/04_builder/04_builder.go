package _4_builder

/**
  用于创建很复杂的对象
*/

// 抽象建造者接口
type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
}

// 规范对象的组件的创建流程
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

// Create方法构造具体的实列
func (d *Director) Create() {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
}

type Build1Instance struct {
	result string
}

func (b *Build1Instance) BuildPart1() {
	b.result += "A"
}

func (b *Build1Instance) BuildPart2() {
	b.result += "B"
}

func (b *Build1Instance) BuildPart3() {
	b.result += "C"
}

func (b *Build1Instance) GetResult() string {
	return b.result
}
