package apierror

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"wing/models/ent"
	"wing/resolver/model"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var errorCodeMap = map[string]model.ErrorCode{}

func errorTypeName(err error) string { return reflect.TypeOf(err).Elem().Name() }

// 用于不明确的错误信息，其中，处理从ent的error转换成接口的错误码和graphql的错误信息
// NOTE 目前ent反馈的数据库的操作错误，一般都会在开发阶段被修正，所以业务逻辑处理，应该不会用到这些错误的转义。
func Transform(ctx context.Context, err error) error {
	var e *gqlerror.Error

	if errors.As(err, &e) {
		return err
	}

	if ent.IsNotFound(err) {
		return nil
	}

	code, ok := errorCodeMap[errorTypeName(err)]
	if !ok {
		code = model.ErrorCodeUnknown
		err = fmt.Errorf("未知错误:%w", err)
	}
	return New(ctx, code, err.Error())
}

// 用于在逻辑处理中产生的明确的错误信息
func New(ctx context.Context, code model.ErrorCode, message string) error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: message,
		Extensions: map[string]interface{}{
			string(model.ErrorCodeKeyErrorCode): code,
		},
	}
}
