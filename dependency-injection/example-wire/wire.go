//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeApplication() (*Application, error) {
	wire.Build(applicationSet)
	return &Application{}, nil
}
