package main

import (
	"carProject/carbrend"
	"carProject/logiccar"
	"fmt"
	"time"
)

func main() {
	var myCar logiccar.Car

	fmt.Println("Твоє авто БВМ")
	myCar = &carbrend.Bmw{Spend: 100, Fuel: 50}

	go fmt.Println(myCar.Rapair(10))
	fmt.Println(myCar.Drive(10))
	fmt.Println(myCar)
	fmt.Println("")

	fmt.Println("Твоє авто Мерседес")
	myCar = &carbrend.Mers{Spend: 100, Fuel: 50}

	fmt.Println(myCar.Rapair(60))
	fmt.Println(myCar.Drive(55))
	fmt.Println(myCar)
	fmt.Println("")

	fmt.Println("Твоє авто Ауді")
	myCar = &carbrend.Audi{Spend: 100, Fuel: 50}

	fmt.Println(myCar.Rapair(145))
	fmt.Println(myCar.Drive(25))
	fmt.Println(myCar)
	time.Sleep(10)
}
