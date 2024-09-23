package main

import (
	"carProject/logiccar"
	"fmt"
)

//type Car interface {
//	rapair(spend int) (int, bool, error)
//	drive(fuel int) (int, error)
//}
//
//type Bmw struct {
//	spend int
//	fuel  int
//}

//	func (b *Bmw) rapair(spend int) (int, bool, error) {
//		spend = b.spend - spend
//
//		if spend < b.spend {
//			fmt.Println("Залишок після ремонту", spend, true)
//			b.spend = spend
//			return spend, true, fmt.Errorf("Ремонт відбувся")
//		} else {
//			fmt.Println("Залишок після ремонту", spend, false, fmt.Errorf("Ремонт не відбувся"))
//			return spend, false, fmt.Errorf("Ремонт не відбувся")
//		}
//
// }
//
//	func (b *Bmw) drive(fuel int) (int, error) {
//		fuel = b.fuel - fuel
//		if fuel < b.fuel {
//			fmt.Println("Залишилось палива", fuel)
//			b.fuel = fuel
//			return fuel, nil
//		} else {
//			return fuel, fmt.Errorf("Палива немає")
//		}
//	}

func main() {
	var myCar logiccar.Car

	fmt.Println("Твоє авто БВМ")
	myCar = &logiccar.LogicBmw{Spend: 100, Fuel: 50}
	fmt.Println(myCar)
	fmt.Println(myCar.Rapair(10))
	fmt.Println(myCar)
	fmt.Println(myCar.Drive(10))
	fmt.Println(myCar)
	fmt.Println("")

	fmt.Println("Твоє авто Мерседес")
	myCar = &logiccar.LogicMers{Spend: 100, Fuel: 50}
	fmt.Println(myCar)
	fmt.Println(myCar.Rapair(60))
	fmt.Println(myCar)
	fmt.Println(myCar.Drive(20))
	fmt.Println(myCar)
	fmt.Println("")

	fmt.Println("Твоє авто Ауді")
	myCar = &logiccar.LogicAudi{Spend: 100, Fuel: 50}
	fmt.Println(myCar)
	fmt.Println(myCar.Rapair(45))
	fmt.Println(myCar)
	fmt.Println(myCar.Drive(25))
	fmt.Println(myCar)

}
