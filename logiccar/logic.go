package logiccar

import (
	"carProject/carbrend"
	"fmt"
)

type Car interface {
	Rapair(Spend int) (int, bool, error)
	Drive(Fuel int) (int, error)
}

type LogicBmw carbrend.Bmw
type LogicMers carbrend.Mers
type LogicAudi carbrend.Audi

func (b *LogicBmw) Rapair(spend int) (int, bool, error) {
	spend = b.Spend - spend

	if spend < b.Spend {
		fmt.Println("Залишок після ремонту", spend, true)
		b.Spend = spend
		return spend, true, fmt.Errorf("Ремонт відбувся")
	} else {
		fmt.Println("Залишок після ремонту", spend, false, fmt.Errorf("Ремонт не відбувся"))
		return spend, false, fmt.Errorf("Ремонт не відбувся")
	}

}

func (b *LogicBmw) Drive(fuel int) (int, error) {
	fuel = b.Fuel - fuel
	if fuel < b.Fuel {
		fmt.Println("Залишилось палива", fuel)
		b.Fuel = fuel
		return fuel, nil
	} else {
		return fuel, fmt.Errorf("Палива немає")
	}
}

func (b *LogicMers) Rapair(spend int) (int, bool, error) {
	spend = b.Spend - spend

	if spend < b.Spend {
		fmt.Println("Залишок після ремонту", spend, true)
		b.Spend = spend
		return spend, true, fmt.Errorf("Ремонт відбувся")
	} else {
		fmt.Println("Залишок після ремонту", spend, false, fmt.Errorf("Ремонт не відбувся"))
		return spend, false, fmt.Errorf("Ремонт не відбувся")
	}

}

func (b *LogicMers) Drive(fuel int) (int, error) {
	fuel = b.Fuel - fuel
	if fuel < b.Fuel {
		fmt.Println("Залишилось палива", fuel)
		b.Fuel = fuel
		return fuel, nil
	} else {
		return fuel, fmt.Errorf("Палива немає")
	}
}

func (b *LogicAudi) Rapair(spend int) (int, bool, error) {
	spend = b.Spend - spend

	if spend < b.Spend {
		fmt.Println("Залишок після ремонту", spend, true)
		b.Spend = spend
		return spend, true, fmt.Errorf("Ремонт відбувся")
	} else {
		fmt.Println("Залишок після ремонту", spend, false, fmt.Errorf("Ремонт не відбувся"))
		return spend, false, fmt.Errorf("Ремонт не відбувся")
	}

}

func (b *LogicAudi) Drive(fuel int) (int, error) {
	fuel = b.Fuel - fuel
	if fuel < b.Fuel {
		fmt.Println("Залишилось палива", fuel)
		b.Fuel = fuel
		return fuel, nil
	} else {
		return fuel, fmt.Errorf("Палива немає")
	}
}
