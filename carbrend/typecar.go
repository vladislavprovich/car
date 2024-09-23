package carbrend

import "fmt"

type Bmw struct {
	Spend int
	Fuel  int
}

type Mers struct {
	Spend int
	Fuel  int
}

type Audi struct {
	Spend int
	Fuel  int
}

func (b *Bmw) Rapair(spend int) (int, bool, error) {
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

func (b *Bmw) Drive(fuel int) (int, error) {
	fuel = b.Fuel - fuel
	if fuel < b.Fuel {
		fmt.Println("Залишилось палива", fuel)
		b.Fuel = fuel
		return fuel, nil
	} else {
		return fuel, fmt.Errorf("Палива немає")
	}
}

func (b *Mers) Rapair(spend int) (int, bool, error) {
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

func (b *Mers) Drive(fuel int) (int, error) {
	fuel = b.Fuel - fuel
	if fuel < b.Fuel {
		fmt.Println("Залишилось палива", fuel)
		b.Fuel = fuel
		return fuel, nil
	} else {
		return fuel, fmt.Errorf("Палива немає")
	}
}

func (b *Audi) Rapair(spend int) (int, bool, error) {
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

func (b *Audi) Drive(fuel int) (int, error) {
	fuel = b.Fuel - fuel
	if fuel < b.Fuel {
		fmt.Println("Залишилось палива", fuel)
		b.Fuel = fuel
		return fuel, nil
	} else {
		return fuel, fmt.Errorf("Палива немає")
	}
}
