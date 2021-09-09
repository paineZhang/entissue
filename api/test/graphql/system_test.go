//go:build ignore
// +build ignore

package apitest

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"wing/api/test/graphql/gen"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	jwtToken string
	ctx      context.Context
	client   *gen.Client
}

func (ins *TestSuite) SetupTest() {
	ins.ctx = context.Background()
	ins.client = &gen.Client{
		Client: clientv2.NewClient(http.DefaultClient, TestAddr,
			func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo,
				res interface{}, next clientv2.RequestInterceptorFunc) error {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ins.jwtToken))

				return next(ctx, req, gqlInfo, res)
			}),
	}
}

func (ins *TestSuite) TearDownTest() {
}

func (ins *TestSuite) TestSystem() {
	_, err := ins.client.GetSystemInfo(ins.ctx)
	assert.NoError(ins.T(), err)

	// fmt.Println(systemInfo.SystemQuery.System.Name)
	// fmt.Println(systemInfo.SystemQuery.System.Name)
	println("ok")

	// if err != nil {
	// 	if handledError, ok := err.(*client.ErrorResponse); ok {
	// 		fmt.Fprintf(os.Stderr, "handled error: %s\n", handledError.Error())
	// 	} else {
	// 		fmt.Fprintf(os.Stderr, "unhandled error: %s\n", err.Error())
	// 	}
	// 	os.Exit(1)
	// }
}

func TestSystem(t *testing.T) {
	suite.Run(t, &TestSuite{})
}
