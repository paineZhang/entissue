package resolver

import (
	"net/http"
	"time"
	"wing/models/ent"
	"wing/resolver/generated"
	"wing/resolver/middleware"
	"wing/resolver/resolvers"

	casbinadapter "wing/adapters/casbin"
	ldapadapter "wing/adapters/ldap"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	apiPath        = "/graphql"
	playgroundPath = "/playground"
	defaultPort    = "9200"
)

func NewHandler(client *ent.Client, ldap *ldapadapter.LdapAdapter, casbin *casbinadapter.CasbinAdapter) http.Handler {
	graphqlApiHandler := graphqlHandler(client, ldap, casbin)
	ginserver := gin.Default()

	ginserver.Use(append([]gin.HandlerFunc{
		gin.Recovery(),
	}, middleware.Middlewares()...)...)

	ginserver.POST(apiPath, graphqlApiHandler)
	ginserver.GET(apiPath, graphqlApiHandler)
	ginserver.GET(playgroundPath, playgroundHandler())

	return ginserver
}

func graphqlHandler(client *ent.Client, ldap *ldapadapter.LdapAdapter, casbin *casbinadapter.CasbinAdapter) gin.HandlerFunc {

	srv := handler.New(generated.NewExecutableSchema(
		generated.Config{Resolvers: &resolvers.Resolver{Ldap: ldap, Casbin: casbin}}))

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	srv.Use(apollotracing.Tracer{})

	// 在QL层的ctx中获取到ent的client，这样在业务层可以直接操作client
	// entgql.Transactioner使得在变更解释操作时，会将client转换成事务的。所以直接使用，不再显示声明事务
	// 注意，这仅仅用在变更操作，查询操作，仍然要使用注入到Resolver的client。有兴趣的话具体参考entgql.Transactioner的实现
	// TODO : Transactioner的InterceptResponse逻辑好像改了，请求操作从context中取出来的不是事务的。逻辑可以统一了。
	srv.Use(entgql.Transactioner{TxOpener: client})

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	srv := playground.Handler("Wing GraphQL playground", apiPath)

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
