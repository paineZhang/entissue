package initialize

import (
	"net/http"
	ldapadapter "wing/adapters/ldap"
	model "wing/models"
	resolver "wing/resolver"

	"github.com/golobby/container/v3"

	casbinadapter "wing/adapters/casbin"
)

// NOTE 系统初始化

func Initialize() http.Handler {

	initContainer()

	client := model.New()

	initSystemInfo(client)

	initResource(client)

	ldap, err := ldapadapter.New()
	if err != nil {
		panic(err)
	}

	var casbin *casbinadapter.CasbinAdapter
	err = container.Resolve(&casbin)
	if err != nil {
		panic(err)
	}

	return resolver.NewHandler(client, ldap, casbin)
}
