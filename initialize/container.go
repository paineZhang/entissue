package initialize

import (
	"github.com/golobby/container/v3"

	casbinadapter "wing/adapters/casbin"
)

func initContainer() {

	err := container.Singleton(func() *casbinadapter.CasbinAdapter {
		casbin, err := casbinadapter.New()
		if err != nil {
			panic(err)
		}
		return casbin
	})

	if err != nil {
		panic(err)
	}
}
