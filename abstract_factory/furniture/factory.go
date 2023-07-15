package furniture

const (
	defaultDiscount = 1.0
)

type basicFactory struct {
	style    Style
	material Material
	discount float64
}

func (f basicFactory) CreateChair() Chair {
	chair := &basicChair{
		legs:     make([]chairLeg, 4),
		armrests: make([]chairArmrest, 2),
	}

	chair.setStyle(f.style)
	chair.setMaterial(f.material)
	chair.setPrice(f.discount)

	return chair
}

func (f basicFactory) CreateTable() Table {
	table := &basicTable{
		legs:   make([]tableLeg, 4),
		things: make(map[string]struct{}),
	}

	table.setStyle(f.style)
	table.setMaterial(f.material)
	table.setPrice(f.discount)
	return table
}

func NewFactory(style Style, material Material, discount ...float64) Factory {
	dis := defaultDiscount
	if len(discount) > 0 {
		dis = discount[0]
	}

	return &basicFactory{
		style:    style,
		material: material,
		discount: dis,
	}
}
