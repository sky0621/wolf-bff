package domain

type ValueObject interface {
	// Validate 何であって何であるべきでないかの定義に即してバリデーション
	Validate() bool
}
