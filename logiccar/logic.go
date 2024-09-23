package logiccar

type Car interface {
	Rapair(Spend int) (int, bool, error)
	Drive(Fuel int) (int, error)
}

// мб можна по іншому
