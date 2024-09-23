package main

import (
	"carProject/carbrend"
	"carProject/logiccar"
	"fmt"
)

func main() {
	var myCar logiccar.Car

	fmt.Println("Твоє авто БВМ")
	myCar = &carbrend.Bmw{Spend: 100, Fuel: 50}
	fmt.Println(myCar)
	fmt.Println(myCar.Rapair(10))
	fmt.Println(myCar.Drive(10))
	fmt.Println(myCar)
	fmt.Println("")

	fmt.Println("Твоє авто Мерседес")
	myCar = &carbrend.Mers{Spend: 100, Fuel: 50}
	fmt.Println(myCar)
	fmt.Println(myCar.Rapair(60))
	fmt.Println(myCar)
	fmt.Println(myCar.Drive(20))
	fmt.Println(myCar)
	fmt.Println("")

	fmt.Println("Твоє авто Ауді")
	myCar = &carbrend.Audi{Spend: 100, Fuel: 50}
	fmt.Println(myCar)
	fmt.Println(myCar.Rapair(45))
	fmt.Println(myCar)
	fmt.Println(myCar.Drive(25))
	fmt.Println(myCar)

}
