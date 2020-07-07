package controller

import "github.com/sky0621/wolf-bff/src/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	wht usecase.WhtInputPort
}

func NewResolver(
	wht usecase.WhtInputPort,
) ResolverRoot {
	return &Resolver{wht: wht}
}
