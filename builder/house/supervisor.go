package house

type basicSupervisor struct {
	WallBuilder   WallBuilder
	WindowBuilder WindowBuilder
	DoorBuilder   DoorBuilder

	wallCount   int
	windowCount int
	doorCount   int
}

func (b basicSupervisor) SetBuilders(builders ...Builder) {
	for i := range builders {
		switch builders[i].(type) {
		case WallBuilder:
			b.WallBuilder = builders[i].(WallBuilder)
		case WindowBuilder:
			b.WindowBuilder = builders[i].(WindowBuilder)
		case DoorBuilder:
			b.DoorBuilder = builders[i].(DoorBuilder)
		case LockBuilder:
			b.DoorBuilder.SetLockBuilder(builders[i].(LockBuilder))
		}
	}
}

func (b basicSupervisor) SetComponentCount(walls, windows, doors int) {
	if walls > 2 {
		b.wallCount = walls
	}

	b.windowCount = windows

	if walls > 0 {
		b.doorCount = doors
	}
}

func (b basicSupervisor) BuildHouse() House {
	house := House{}

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

		wallCount:   5,
		windowCount: 4,
		doorCount:   1,
	}
}
