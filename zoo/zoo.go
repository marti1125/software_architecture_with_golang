package zoo

import "fmt"

type LatLog struct {
	Lat  float64
	Long float64
}

type Animal interface {
	GetLocation() LatLog
	SetLocation(LatLog)
	CanFly() bool
	Speak() string
	GetName() string
}

func MakeThemSing(animals []Animal) {
	for _, animal := range animals {
		fmt.Println(animal.GetName() + " says " + animal.Speak())
	}
}
