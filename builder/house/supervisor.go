package house

type basicSupervisor struct {
	WallBuilder   WallBuilder
	WindowBuilder WindowBuilder
	DoorBuilder   DoorBuilder
	RoofBuilder   RoofBuilder

	wallCount   int
	windowCount int
	doorCount   int
}

func (b basicSupervisor) SetBuilders(builders ...ComponentBuilder) {
	for i := range builders {
		switch builders[i].(type) {
		case WallBuilder:
			b.WallBuilder = builders[i].(WallBuilder)
		case WindowBuilder:
			b.WindowBuilder = builders[i].(WindowBuilder)
		case DoorBuilder:
			b.DoorBuilder = builders[i].(DoorBuilder)
		case RoofBuilder:
			b.RoofBuilder = builders[i].(RoofBuilder)
		case LockBuilder:
			b.DoorBuilder.SetLockBuilder(builders[i].(LockBuilder))
		}
	}
}

// SetComponentCount sets the number of walls, windows and doors to be built.
// The number of walls must be at least 3.
// The number of windows and doors must be less than or equal to the number of walls * 2.
func (b basicSupervisor) SetComponentCount(walls, windows, doors int) {
	// there must be at least 3 walls
	if walls > 2 {
		b.wallCount = walls
	}

	// there can not be more windows than walls * 2 and there must be at least one window
	if windows <= b.wallCount*2 && windows > 0 {
		b.windowCount = windows
	}

	// there can not be more doors than walls * 2 and there must be at least one door
	if doors <= b.wallCount*2 && doors > 0 {
		b.doorCount = doors
	}
}

func (b basicSupervisor) BuildHouse() House {
	house := House{
		Roof: b.RoofBuilder.Build(),
	}

	house.Walls = make([]Wall, b.wallCount)
	for i := 0; i < b.wallCount; i++ {
		house.Walls[i] = b.WallBuilder.Build()
	}

	house.Windows = make([]Window, b.windowCount)
	for i := 0; i < b.windowCount; i++ {
		house.Windows[i] = b.WindowBuilder.Build()
	}

	house.Doors = make([]Door, b.doorCount)
	for i := 0; i < b.doorCount; i++ {
		house.Doors[i] = b.DoorBuilder.Build()
	}

	return house
}

func NewSupervisor() Supervisor {
	return &basicSupervisor{
		WallBuilder:   NewWallBuilder(),
		WindowBuilder: NewWindowBuilder(),
		DoorBuilder:   NewDoorBuilder(),
		RoofBuilder:   NewRoofBuilder(),

		wallCount:   4,
		windowCount: 2,
		doorCount:   1,
	}
}
