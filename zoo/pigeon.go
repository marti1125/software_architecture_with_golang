package zoo

type Pigeon struct {
	Name     string
	location LatLog
}

func (p *Pigeon) GetLocation() LatLog {
	return p.location
}

func (p *Pigeon) SetLocation(loc LatLog) {
	p.location = loc
}

func (p *Pigeon) CanFly() bool {
	return false
}

func (p *Pigeon) Speak() string {
	return "hoot"
}

func (p *Pigeon) GetName() string {
	return p.Name
}
