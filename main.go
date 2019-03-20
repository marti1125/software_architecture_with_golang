package main

//import "./pigeon"
import "./zoo"

/*type Animal struct {
	Name   string
	canFly bool
}

type Person struct {
	Name string
	Age  int
}

func (p Person) canVote() bool {
	return p.Age > 18
}

func (p *Person) Grow() {
	p.Age++
}

func (p Person) DoesNotGrow() {
	p.Age++
}*/

func main() {
	/*anAnimal := Animal{Name: "Lion", canFly: false}
	//fmt.Printf("hello, world\n")
	fmt.Println(anAnimal.Name)
	aLionPtr := &anAnimal
	//println(aLionPtr.age) error...
	println(aLionPtr)
	aPerson := Person{Name: "Willy", Age: 30}
	println(aPerson.Age)
	aPerson.Grow()
	println(aPerson.canVote())
	println(aPerson.Age)
	ptr := &aPerson
	ptr.DoesNotGrow()
	println(ptr.Age)

	p := pigeon.Pigeon{Name: "Tweenty"}
	p.SetFeatherLength(10)
	println(p.Name)
	println(p.GetFeatherLength())*/

	// ZOO

	var myZoo []zoo.Animal

	Leo := zoo.Lion{Name: "Leo"}

	LeoLoc := zoo.LatLog{10.3, 10.4}

	Leo.SetLocation(LeoLoc)

	myZoo = append(myZoo, &Leo)

	Tweety := zoo.Pigeon{Name: "Tweety"}

	TwLoc := zoo.LatLog{130.3, 1033.4}

	Tweety.SetLocation(TwLoc)

	myZoo = append(myZoo, &Tweety)

	zoo.MakeThemSing(myZoo)

}
