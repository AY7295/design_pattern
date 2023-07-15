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

type House struct {
	Windows []Window
	Walls   []Wall
	Doors   []Door
}

func (h House) Information() map[string]any {
	return map[string]any{
		"windows": h.Windows,
		"walls":   h.Walls,
		"doors":   h.Doors,
	}
}

type Builder interface {
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
	Builder
	SetThickness(thickness int)
	Build() Wall
}

type WindowBuilder interface {
	Builder
	Build() Window
}

type DoorBuilder interface {
	Builder
	SetThickness(thickness int)
	SetLockBuilder(lockBuilder LockBuilder)
	Build() Door
}

type Supervisor interface {
	SetBuilders(builders ...Builder)
	SetComponentCount(walls, windows, doors int)
	BuildHouse() House
}
