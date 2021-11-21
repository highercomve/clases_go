package main

import (
	"fmt"

	"github.com/highercomve/clases_go/utils"
)

func init() {
	fmt.Println("Inicio")
}

func main() {
	fmt.Println("Comienza main")

	carro := utils.NewCarro(utils.Position{X: 0, Y: 0})
	bici := utils.NewBicicleta(utils.Position{X: 0, Y: 0})

	fmt.Printf("carro ---\t %+v\t%p\n", *carro, carro)
	fmt.Printf("bici ---\t %+v\t%p\n", *bici, bici)

	utils.MoverVehiculo(carro)
	utils.MoverVehiculo(bici)

	fmt.Printf("carro ---\t %+v\t%p\n", *carro, carro)
	fmt.Printf("bici ---\t %+v\t%p\n", *bici, bici)

	utils.CojeVehiculo(carro)
	utils.CojeVehiculo(bici)

	fmt.Printf("carro ---\t %+v\t%p\n", *carro, carro)
	fmt.Printf("bici ---\t %+v\t%p\n", *bici, bici)
}
