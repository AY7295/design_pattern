package house

type basicRoofBuilder struct {
	material string
	color    string
	size     int
	shape    string
}

func (b *basicRoofBuilder) SetMaterial(material string) {
	b.material = material
}

func (b *basicRoofBuilder) SetColor(color string) {
	b.color = color
}

func (b *basicRoofBuilder) SetSize(size int) {
	b.size = size
}

func (b *basicRoofBuilder) SetShape(shape string) {
	b.shape = shape
}

func (b *basicRoofBuilder) Build() Roof {
	return Roof{
		Material: b.material,
		Color:    b.color,
		Size:     b.size,
		Shape:    b.shape,
	}
}

func CustomRoofBuilder(material, color string, size int, shape string) RoofBuilder {
	return &basicRoofBuilder{
		material: material,
		color:    color,
		size:     size,
		shape:    shape,
	}
}

func NewRoofBuilder() RoofBuilder {
	return &basicRoofBuilder{
		material: "wood",
		color:    "white",
		size:     20,
		shape:    "flat",
	}
}
