package house

type Window struct {
	Material string
	Color    string
	Size     int
}

type Wall struct {
	Material  string
	Color     string
	Size      int
	Thickness int
}

type Door struct {
	Material  string
	Color     string
	Size      int
	Thickness int
	Lock      Lock
}

type Lock struct {
	Material string
	Color    string
	Size     int
}

type Roof struct {
	Material string
	Color    string
	Size     int
	Shape    string
}

type House struct {
	Windows []Window
	Walls   []Wall
	Doors   []Door
	Roof    Roof
}

func (h House) Information() map[string]any {
	return map[string]any{
		"windows": h.Windows,
		"walls":   h.Walls,
		"doors":   h.Doors,
		"roof":    h.Roof,
	}
}

type ComponentBuilder interface {
	SetMaterial(material string)
	SetColor(color string)
	SetSize(size int)
}

type LockBuilder interface {
	SetMaterial(material string)
	SetColor(color string)
	SetSize(size int)
	Build() Lock
}

type WallBuilder interface {
	ComponentBuilder
	SetThickness(thickness int)
	Build() Wall
}

type WindowBuilder interface {
	ComponentBuilder
	Build() Window
}

type DoorBuilder interface {
	ComponentBuilder
	SetThickness(thickness int)
	SetLockBuilder(lockBuilder LockBuilder)
	Build() Door
}

type RoofBuilder interface {
	ComponentBuilder
	SetShape(shape string)
	Build() Roof
}

type Supervisor interface {
	SetBuilders(builders ...ComponentBuilder)
	SetComponentCount(walls, windows, doors int)
	BuildHouse() House
}
