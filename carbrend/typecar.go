package carbrend

import "fmt"

type Bmw struct {
	Spend uint
	Fuel  uint
}

type Mers struct {
	Spend uint
	Fuel  uint
}

type Audi struct {
	Spend uint
	Fuel  uint
}

func (b *Bmw) Rapair(spend uint) (uint, bool, error) {
	spend = b.Spend - spend

	if spend < b.Spend {
		fmt.Print("Залишок після ремонту ")
		b.Spend = spend
		return spend, true, fmt.Errorf(",Ремонт відбувся")
	} else {
		fmt.Print("Залишок після ремонту ")
		return 0, false, fmt.Errorf(",Ремонт не відбувся")
	}
}

func (b *Bmw) Drive(fuel uint) (uint, error) {
	fuel = b.Fuel - fuel
	if fuel < b.Fuel {
		fmt.Print("Залишилось палива ")
		b.Fuel = fuel
		return fuel, fmt.Errorf(",палива вистачить")
	} else {
		b.Fuel = 0
		return 0, fmt.Errorf("Палива немає")
	}
}

func (b *Mers) Rapair(spend uint) (uint, bool, error) {
	spend = b.Spend - spend

	if spend < b.Spend {
		fmt.Print("Залишок після ремонту ")
		b.Spend = spend
		return spend, true, fmt.Errorf(",Ремонт відбувся")
	} else {
		fmt.Print("Залишок після ремонту ")
		return 0, false, fmt.Errorf(",Ремонт не відбувся")
	}

}

func (b *Mers) Drive(fuel uint) (uint, error) {
	fuel = b.Fuel - fuel
	if fuel < b.Fuel {
		fmt.Print("Залишилось палива ")
		b.Fuel = fuel
		return fuel, fmt.Errorf(",палива вистачить")
	} else {
		b.Fuel = 0
		return 0, fmt.Errorf("Палива немає")
	}
}

func (b *Audi) Rapair(spend uint) (uint, bool, error) {
	spend = b.Spend - spend

	if spend < b.Spend {
		fmt.Print("Залишок після ремонту ")
		b.Spend = spend
		return spend, true, fmt.Errorf(",Ремонт відбувся")
	} else {
		fmt.Print("Залишок після ремонту ")
		return 0, false, fmt.Errorf(",Ремонт не відбувся")
	}

}

func (b *Audi) Drive(fuel uint) (uint, error) {
	fuel = b.Fuel - fuel
	if fuel < b.Fuel {
		fmt.Print("Залишилось палива ")
		b.Fuel = fuel
		return fuel, fmt.Errorf(",палива вистачить")
	} else {
		b.Fuel = 0
		return 0, fmt.Errorf("Палива немає")
	}
}
