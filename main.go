package main

import "fmt"

type Car interface {
	rapair(spend int) (int, bool, error)
	drive(fuel int) (int, error)
}

type Bmw struct {
	spend int
	fuel  int
}

func (b *Bmw) rapair(spend int) (int, bool, error) {
	spend = b.spend - spend

	if spend < b.spend {
		fmt.Println("Залишок після ремонту", spend, true)
		b.spend = spend
		return spend, true, fmt.Errorf("Ремонт відбувся")
	} else {
		fmt.Println("Залишок після ремонту", spend, false, fmt.Errorf("Ремонт не відбувся"))
		return spend, false, fmt.Errorf("Ремонт не відбувся")
	}

}

func (b *Bmw) drive(fuel int) (int, error) {
	fuel = b.fuel - fuel
	if fuel < b.fuel {
		fmt.Println("Залишилось палива", fuel)
		b.fuel = fuel
		return fuel, nil
	} else {
		return fuel, fmt.Errorf("Палива немає")
	}
}

func main() {
	var myCar Car

	myCar = &Bmw{spend: 100, fuel: 50}
	fmt.Println(myCar)
	fmt.Println(myCar.rapair(10))
	fmt.Println(myCar)
	fmt.Println(myCar.drive(10))
	fmt.Println(myCar)
}
