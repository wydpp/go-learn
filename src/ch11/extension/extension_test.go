package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(o string) {
	p.Speak()
	fmt.Println(" ", o)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

/*func (d *Dog) SpeakTo(o string){
	d.p.SpeakTo(o)
}*/

func TestDog(t *testing.T) {
	var dog = new(Dog)
	dog.Speak()
	dog.SpeakTo("Chao")
}
