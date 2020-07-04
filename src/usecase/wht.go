package usecase

type Wht interface {
	CreateWht()
}

type wht struct {
}

func NewWht() Wht {
	return &wht{}
}

func (w *wht) CreateWht() {

}
