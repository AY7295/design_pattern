package furniture

type basicTable struct {
	legs     []tableLeg
	top      tableTop
	price    float64
	things   map[string]struct{}
	style    Style
	material Material
}

func (t *basicTable) getStyle() Style {
	return t.style
}

func (t *basicTable) getMaterial() Material {
	return t.material
}

func (t *basicTable) Information() map[string]any {
	return map[string]any{
		"legs":   t.legs,
		"top":    t.top,
		"price":  t.price,
		"things": t.things,
	}
}

type tableLeg struct {
	style    Style
	material Material
}

type tableTop struct {
	style    Style
	material Material
}

func (t *basicTable) GetPrice() float64 {
	return t.price
}

func (t *basicTable) setPrice(discount float64) {
	t.price = float64(t.style) * float64(t.material) * discount
}

func (t *basicTable) setStyle(style Style) {
	t.style = style
	t.top.style = style
	for i := 0; i < len(t.legs); i++ {
		t.legs[i].style = style
	}
}

func (t *basicTable) setMaterial(material Material) {
	t.material = material
	t.top.material = material
	for i := 0; i < len(t.legs); i++ {
		t.legs[i].material = material
	}
}

func (t *basicTable) Place(thing string) {
	t.things[thing] = struct{}{}
}

func (t *basicTable) Remove(thing string) {
	delete(t.things, thing)
}
