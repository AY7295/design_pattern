package furniture

type basicChair struct {
	legs     []chairLeg
	armrests []chairArmrest
	back     chairBack
	price    float64
	isSet    bool
	style    Style
	material Material
}

func (c *basicChair) getStyle() Style {
	return c.style
}

func (c *basicChair) getMaterial() Material {
	return c.material
}

func (c *basicChair) Information() map[string]any {
	return map[string]any{
		"legs":     c.legs,
		"armrests": c.armrests,
		"back":     c.back,
		"price":    c.price,
	}
}

// chairLeg 座椅腿
type chairLeg struct {
	style    Style
	material Material
}

// chairBack 座椅靠背
type chairBack struct {
	style    Style
	material Material
}

// chairArmrest 座椅扶手
type chairArmrest struct {
	style    Style
	material Material
}

func (c *basicChair) GetPrice() float64 {
	return c.price
}

func (c *basicChair) setPrice(discount float64) {
	c.price = float64(c.style) * float64(c.material) * discount
}

func (c *basicChair) SetOn() {
	c.isSet = true
}

func (c *basicChair) SetOff() {
	c.isSet = false
}

func (c *basicChair) setStyle(style Style) {
	c.style = style
	c.back.style = style
	for i := 0; i < len(c.legs); i++ {
		c.legs[i].style = style
	}
	for i := 0; i < len(c.armrests); i++ {
		c.armrests[i].style = style
	}
}

func (c *basicChair) setMaterial(material Material) {
	c.material = material
	c.back.material = material
	for i := 0; i < len(c.legs); i++ {
		c.legs[i].material = material
	}
	for i := 0; i < len(c.armrests); i++ {
		c.armrests[i].material = material
	}
}
