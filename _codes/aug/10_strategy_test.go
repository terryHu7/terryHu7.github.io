package aug

import (
	"fmt"
	"testing"
)

type Vehicle interface {
	Go()
}

type Car struct {
}

func (r *Car) Go() {
	fmt.Println("use car")
}

type Bicycle struct {
}

func (r *Bicycle) Go() {
	fmt.Println("use Bicycle")
}

type Traveler struct {
	Vehicle
}

func (r *Traveler) SetVehicle(i Vehicle) {
	r.Vehicle = i
}

func (r *Traveler) Go() {
	r.Vehicle.Go()
}

func TestStrategy(t *testing.T) {
	traveler := Traveler{}

	traveler.Vehicle = &Car{} // traveler.SetVehicle(&Car{})
	traveler.Go()

	traveler.Vehicle = &Bicycle{} // traveler.SetVehicle(&Bicycle{})
	traveler.Go()
}
