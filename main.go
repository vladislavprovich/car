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

	if res, conflict, err := myCar.Repair(10); err != nil {
		fmt.Println("", res, conflict, err)
	}
	if res, err := myCar.Drive(10); err != nil {
		fmt.Println("", res, err)
	}
	fmt.Println(myCar)
	fmt.Println("")

	fmt.Println("Твоє авто Мерседес")
	myCar = &carbrend.Mers{Spend: 100, Fuel: 50}

	if res, conflict, err := myCar.Repair(80); err != nil {
		fmt.Println("", res, conflict, err)
	}
	if res, err := myCar.Drive(60); err != nil {
		fmt.Println("", res, err)
	}
	fmt.Println(myCar)
	fmt.Println("")

	fmt.Println("Твоє авто Ауді")
	myCar = &carbrend.Audi{Spend: 100, Fuel: 50}

	if res, conflict, err := myCar.Repair(140); err != nil {
		fmt.Println("", res, conflict, err)
	}
	if res, err := myCar.Drive(25); err != nil {
		fmt.Println("", res, err)
	}
	fmt.Println(myCar)
	time.Sleep(10)
}
