package domain

type ValueObject interface {
	Validate() bool
}
