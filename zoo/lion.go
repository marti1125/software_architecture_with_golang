package zoo

type Lion struct {
	Name       string
	ManeLegnth int
	location   LatLog
}

func (lion *Lion) GetLocation() LatLog {
	return lion.location
}

func (lion *Lion) SetLocation(loc LatLog) {
	lion.location = loc
}

func (lion *Lion) CanFly() bool {
	return false
}

func (lion *Lion) Speak() string {
	return "roar"
}

func (lion *Lion) GetManeLength() int {
	return lion.ManeLegnth
}

func (lion *Lion) GetName() string {
	return lion.Name
}
