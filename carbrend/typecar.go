package carbrend

type car interface {
	rapair(spend int) (int, bool, error)
	drive(fuel int) (int, error)
}

type Bmw struct {
}

type Mers struct {
}

type Audi struct {
}
