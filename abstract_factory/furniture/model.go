package furniture

type Material int

const (
	Wood Material = (iota + 1) * 10
	Plastic
)

type Style int

const (
	ModernStyle Style = (iota + 1) * 10
	VictorianStyle
)

type Factory interface {
	CreateChair() Chair
	CreateTable() Table
}

type Furniture interface {
	Information() map[string]any
	GetPrice() float64
	setPrice(discount float64)
	setStyle(style Style)
	setMaterial(material Material)
	getStyle() Style
	getMaterial() Material
}

type Chair interface {
	Furniture
	SetOn()
	SetOff()
}

type Table interface {
	Furniture
	Place(thing string)
	Remove(thing string)
}
