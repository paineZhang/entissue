package resolvers

import (
	casbinadapter "wing/adapters/casbin"
	ldapadapter "wing/adapters/ldap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// Model *ent.Client
	Ldap   *ldapadapter.LdapAdapter
	Casbin *casbinadapter.CasbinAdapter
}
