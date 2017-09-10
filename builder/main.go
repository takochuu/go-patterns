package main

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Action
}

type Action interface {
	Drive() error
	Stop() error
}

type CompleteCar struct{}

func newCompleteCar() *CompleteCar {
	return &CompleteCar{}
}
func (c *CompleteCar) Drive() error {
	fmt.Println("Car Driving")
	return nil
}
func (c *CompleteCar) Stop() error {
	fmt.Println("Car Stopping")
	return nil
}

type Car struct {
	color Color
	wheel Wheels
	speed Speed
}

func (c *Car) Color(cl Color) Builder {
	c.color = cl
	return c
}

func (c *Car) Wheels(wh Wheels) Builder {
	c.wheel = wh
	return c
}

func (c *Car) TopSpeed(sp Speed) Builder {
	c.speed = sp
	return c
}

func (c *Car) Build() Action {
	return newCompleteCar()
}

func NewCar() *Car {
	return &Car{}
}

func main() {
	car := NewCar().Color(BlueColor).Wheels(SportsWheels).TopSpeed(KPH).Build()
	car.Drive()
}
