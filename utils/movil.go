package utils

import (
	"math"
)

type Carro struct {
	Position Position
	Type     string
	Wheels   int
}

type Position struct {
	X int
	Y int
}

func NewCarro(p Position) *Carro {
	return &Carro{
		Position: p,
		Type:     "Car",
		Wheels:   4,
	}
}

func (m *Carro) Moverse(d float64, speed int, time int) Position {
	radians := d * math.Pi / 180
	m.Position.X = m.Position.X + (int(math.Sin(radians)) * speed * time)
	m.Position.Y = m.Position.Y + (int(math.Cos(radians)) * speed * time)

	return m.Position
}

func (m *Carro) ResetPosition() {
	m.Position.X = 0
	m.Position.Y = 0
}
