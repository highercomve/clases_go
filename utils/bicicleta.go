package utils

import (
	"math"
)

type Bicicleta struct {
	Position Position
	Type     string
	Wheels   int
}

func NewBicicleta(p Position) *Bicicleta {
	return &Bicicleta{
		Position: p,
		Type:     "Bicicleta",
		Wheels:   2,
	}
}

func (m *Bicicleta) Moverse(d float64, speed int, time int) Position {
	radians := d * math.Pi / 180
	m.Position.X = m.Position.X + (int(math.Sin(radians)) * speed * time)
	m.Position.Y = m.Position.Y + (int(math.Cos(radians)) * speed * time)

	return m.Position
}

func (m *Bicicleta) ResetPosition() {
	m.Position.X = 0
	m.Position.Y = 0
}
