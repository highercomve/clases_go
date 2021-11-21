package utils

import "fmt"

type Vehiculo interface {
	Moverse(d float64, speed int, time int) Position
	ResetPosition()
}

func MoverVehiculo(v Vehiculo) {
	v.Moverse(90, 20, 1)
}

func CojeVehiculo(c Vehiculo) {
	fmt.Printf("CojeCarro \t %+v\t%p\n", c, &c)

	c.ResetPosition()

	fmt.Printf("CojeCarro \t %+v\t%p\n", c, &c)
}
