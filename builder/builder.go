package builder

type Product struct {
	Party1 string
	Party2 string
	Party3 string
}

type Builder interface {
	BuildParty1()
	BuildParty2()
	BuildParty3()
	GetResult() Product
}

type ConCreateBuild struct {
	product Product
}

func (b *ConCreateBuild) BuildParty1() {
	b.product.Party1 = "产品1"
}

func (b *ConCreateBuild) BuildParty2() {
	b.product.Party2 = "产品2"
}

func (b *ConCreateBuild) BuildParty3() {
	b.product.Party3 = "产品3"
}

func (b *ConCreateBuild) GetResult() Product {
	return b.product
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.BuildParty1()
	d.builder.BuildParty2()
	d.builder.BuildParty3()
}

func NewDirector(builder Builder) *Director {
	return &Director{builder}
}
