package logiccar

type Car interface {
	Rapair(Spend uint) (uint, bool, error)
	Drive(Fuel uint) (uint, error)
}

// мб можна по іншому
