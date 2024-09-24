package logiccar

type Car interface {
	Repair(Spend uint) (uint, bool, error)
	Drive(Fuel uint) (uint, error)
}
