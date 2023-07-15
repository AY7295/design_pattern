package house

type basicWindowBuilder struct {
	material string
	color    string
	size     int
}

func (b *basicWindowBuilder) SetMaterial(material string) {
	b.material = material
}

func (b *basicWindowBuilder) SetColor(color string) {
	b.color = color
}

func (b *basicWindowBuilder) SetSize(size int) {
	b.size = size
}

func (b *basicWindowBuilder) Build() Window {
	return Window{
		Material: b.material,
		Color:    b.color,
		Size:     b.size,
	}
}

func NewWindowBuilder() WindowBuilder {
	return &basicWindowBuilder{
		material: "wood",
		color:    "white",
		size:     10,
	}
}

func CustomWindowBuilder(material, color string, size int) WindowBuilder {
	return &basicWindowBuilder{
		material: material,
		color:    color,
		size:     size,
	}
}
