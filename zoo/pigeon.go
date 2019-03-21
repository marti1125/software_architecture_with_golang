package zoo

type Bird struct {
	classification string
}

type Pigeon struct {
	Bird
	Name     string
	location LatLog
}

func (b *Bird) GetClassification() string {
	return b.classification
}

func (b *Bird) SetClassification(cla string) {
	b.classification = cla
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
