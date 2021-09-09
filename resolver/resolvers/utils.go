package resolvers

import (
	"github.com/golobby/container/v3"
)

func containerResolve(abstraction interface{}) {

	err := container.Resolve(abstraction)
	if err != nil {
		panic(err)
	}
}
