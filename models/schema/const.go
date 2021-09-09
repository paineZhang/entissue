package schema

import "entgo.io/ent/dialect"

type (
	UserIdContextKey struct{}
)

var (
	MySqlTimeSchemaType = map[string]string{dialect.MySQL: "datetime"}
	MySqlDateSchemaType = map[string]string{dialect.MySQL: "date"}

	TimeSchemaType = MySqlTimeSchemaType
	DateSchemaType = MySqlDateSchemaType
)

const ()
