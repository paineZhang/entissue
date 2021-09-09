package ldapadapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLdap(t *testing.T) {

	ldap, err := New()
	assert.NoError(t, err)

	defer ldap.Uninitialize()

	person := &PersonModel{
		AccountName: "zhangl",
		Password:    "zhangl",
		Mail:        "zhangl@meixing.com",
		DisplayName: "张璐",
		FamilyName:  "张",
		GivenName:   "璐",
	}

	err = ldap.Add(person)
	assert.NoError(t, err)

	pr, err := ldap.Get("zhangl")
	assert.NoError(t, err)
	pt := *person
	pt.Password = ""
	assert.EqualValues(t, pr, &pt)

	// 重复添加
	err = ldap.Add(person)
	assert.Error(t, err)

	err = ldap.Authentication("zhangl", "zhangl")
	assert.NoError(t, err)

	// 认证不存在或密码错误
	err = ldap.Authentication("zhangl", "zhangl1")
	assert.Error(t, err)
	err = ldap.Authentication("zhangl1", "zhangl")
	assert.Error(t, err)

	bRes, err := ldap.Find("zhangl")
	assert.True(t, bRes)
	assert.NoError(t, err)
	bRes, err = ldap.Find("zhangl1")
	assert.False(t, bRes)
	assert.NoError(t, err)

	err = ldap.ResetPassword("zhangl", "zhangl1")
	assert.NoError(t, err)
	err = ldap.Authentication("zhangl", "zhangl")
	assert.Error(t, err)
	err = ldap.Authentication("zhangl", "zhangl1")
	assert.NoError(t, err)

	err = ldap.Delete("zhangl")
	assert.NoError(t, err)
	err = ldap.Delete("zhangl1")
	assert.Error(t, err)
}
