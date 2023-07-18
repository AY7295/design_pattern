package house

type basicDoorBuilder struct {
	material  string
	color     string
	size      int
	thickness int

	lockBuilder LockBuilder
}

func (b *basicDoorBuilder) SetLockBuilder(lockBuilder LockBuilder) {
	b.lockBuilder = lockBuilder
}

func (b *basicDoorBuilder) SetMaterial(material string) {
	b.material = material
}

func (b *basicDoorBuilder) SetColor(color string) {
	b.color = color
}

func (b *basicDoorBuilder) SetSize(size int) {
	b.size = size
}

func (b *basicDoorBuilder) SetThickness(thickness int) {
	b.thickness = thickness
}

func (b *basicDoorBuilder) Build() Door {
	return Door{
		Material:  b.material,
		Color:     b.color,
		Size:      b.size,
		Thickness: b.thickness,
		Lock:      b.lockBuilder.Build(),
	}
}

type basicLockBuilder struct {
	material string
	color    string
	size     int
}

func (b *basicLockBuilder) SetMaterial(material string) {
	b.material = material
}

func (b *basicLockBuilder) SetColor(color string) {
	b.color = color
}

func (b *basicLockBuilder) SetSize(size int) {
	b.size = size
}

func (b *basicLockBuilder) Build() Lock {
	return Lock{
		Material: b.material,
		Color:    b.color,
		Size:     b.size,
	}
}

func NewLockBuilder() LockBuilder {
	return &basicLockBuilder{
		material: "steel",
		color:    "silver",
		size:     10,
	}
}

func CustomLockBuilder(material, color string, size int) LockBuilder {
	return &basicLockBuilder{
		material: material,
		color:    color,
		size:     size,
	}
}

func NewDoorBuilder() DoorBuilder {
	return &basicDoorBuilder{
		material:    "wood",
		color:       "brown",
		size:        100,
		thickness:   10,
		lockBuilder: NewLockBuilder(),
	}

}

func CustomDoorBuilder(material, color string, size, thickness int, lockBuilder ...LockBuilder) DoorBuilder {
	doorBuilder := &basicDoorBuilder{
		material:  material,
		color:     color,
		size:      size,
		thickness: thickness,
	}
	if len(lockBuilder) != 0 && lockBuilder[0] != nil {
		doorBuilder.lockBuilder = lockBuilder[0]
	} else {
		doorBuilder.lockBuilder = CustomLockBuilder(material, color, size/10)
	}

	return doorBuilder
}
