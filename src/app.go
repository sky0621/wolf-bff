package wht

type App interface {
}

type app struct {
}

func NewApp() App {
	return &app{}
}
