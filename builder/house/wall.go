package house

type basicWallBuilder struct {
	material  string
	color     string
	size      int
	thickness int
}

func (b *basicWallBuilder) SetMaterial(material string) {
	b.material = material
}

func (b *basicWallBuilder) SetColor(color string) {
	b.color = color
}

func (b *basicWallBuilder) SetSize(size int) {
	b.size = size
}

func (b *basicWallBuilder) SetThickness(thickness int) {
	b.thickness = thickness
}

func (b *basicWallBuilder) Build() Wall {
	return Wall{
		Material:  b.material,
		Color:     b.color,
		Size:      b.size,
		Thickness: b.thickness,
	}
}

func CustomWallBuilder(material, color string, size, thickness int) WallBuilder {
	return &basicWallBuilder{
		material:  material,
		color:     color,
		size:      size,
		thickness: thickness,
	}
}

func NewWallBuilder() WallBuilder {
	return &basicWallBuilder{
		material:  "wood",
		color:     "white",
		size:      10,
		thickness: 1,
	}
}
