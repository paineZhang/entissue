package model

import (
	"context"
	"fmt"

	"wing/models/ent"
	"wing/models/ent/migrate"

	"database/sql"
	"database/sql/driver"

	_ "wing/models/ent/runtime"

	"contrib.go.opencensus.io/integrations/ocsql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-sql-driver/mysql"
)

type connector struct {
	dsn string
}

func (c connector) Connect(context.Context) (driver.Conn, error) {
	return c.Driver().Open(c.dsn)
}

func (connector) Driver() driver.Driver {
	return ocsql.Wrap(
		mysql.MySQLDriver{},
		ocsql.WithAllTraceOptions(),
		ocsql.WithRowsClose(false),
		ocsql.WithRowsNext(false),
		ocsql.WithDisableErrSkip(true),
	)
}

const (
	dbAddressDefault      = "127.0.0.1"
	dbPortDefault         = "3306"
	dbUserNameDefault     = "wing"
	dbUserPasswordDefault = "wing"
	dbDatabaseNameDefault = "wing"
)

func New() *ent.Client {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		dbUserNameDefault, dbUserPasswordDefault, dbAddressDefault, dbPortDefault, dbDatabaseNameDefault)

	drv := entsql.OpenDB(dialect.MySQL, sql.OpenDB(connector{dsn}))

	// TODO 使用环境变量来处理开发环境、测试环境、生产环境问题。
	client := ent.NewClient(ent.Driver(drv), ent.Debug())

	// NOTE
	// migrate.WithGlobalUniqueID(true),该选项用于graphql对node接口的支持，要求所有对象全局ID唯一。
	// 需要更加注意的是，WithDropIndex，尤其是WithDropColumn的选项。如果需要产生手动迁移过程，那么需要再
	// 当前自动迁移前，完成数据的处理。
	if err := client.Schema.Create(context.Background(),
		migrate.WithFixture(true),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
		panic(err)
	}

	return client
}
