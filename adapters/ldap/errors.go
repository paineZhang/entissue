package ldapadapter

import (
	"fmt"

	ad "github.com/go-ldap/ldap/v3"
)

// 封装的ldap v3，在错误里只有code，没有message。封装一下。

type LdapError struct {
	Code    uint16
	Message string
}

// 实现标准库Error接口
func (e *LdapError) Error() string {
	return fmt.Sprintf("ldap error: %d => %s", e.Code, e.Message)
}

func NewError(err error) *LdapError {
	var code uint16
	e, ok := err.(*ad.Error)
	if ok {
		code = e.ResultCode
	} else {
		code = ad.LDAPResultOther
	}

	return &LdapError{Code: code, Message: ad.LDAPResultCodeMap[code]}
}

func NewErrorWithMessage(code uint16, message string) *LdapError {

	return &LdapError{Code: code, Message: message}
}

func IsAuthUnKnown(err error) bool {
	e, ok := err.(*LdapError)
	if ok {
		return e.Code == ad.LDAPResultAuthUnknown
	}
	return false
}
