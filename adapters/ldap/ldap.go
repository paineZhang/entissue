package ldapadapter

import (
	"github.com/go-ldap/ldap/v3"
)

const (
	// TODO 取环境变量的。否则使用缺省的
	ldapAdderDefault = "ldap://127.0.0.1:389"

	rootDN = "dc=meixing,dc=com"

	bindUser     = "cn=admin," + rootDN
	bindPassword = "wing"

	ouBaseDN = "ou=people," + rootDN
)

type LdapAdapter struct {
	conn *ldap.Conn
}

func New() (l *LdapAdapter, err error) {
	l = &LdapAdapter{}

	l.conn, err = ldap.DialURL(ldapAdderDefault)
	if err != nil {
		return nil, NewError(err)
	}

	err = l.conn.Bind(bindUser, bindPassword)
	if err != nil {
		return nil, NewError(err)
	}

	return l, nil
}

func (ins *LdapAdapter) Uninitialize() {
	ins.conn.Close()
}

func getDN(name string) string {
	return "uid=" + name + "," + ouBaseDN
}

func (ins *LdapAdapter) searchOUAccountDN(accountName string) (*ldap.SearchResult, error) {
	res, err := ins.conn.Search(ldap.NewSearchRequest(
		ouBaseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=inetOrgPerson)(uid="+accountName+"))",
		[]string{"dn"}, // A list attributes to retrieve
		nil,
	))

	if err != nil {
		return nil, NewError(err)
	}

	if len(res.Entries) == 0 {
		return nil, nil
	} else if len(res.Entries) > 1 {
		return nil, NewErrorWithMessage(ldap.LDAPResultOther, "存在多个相同账户名")
	}

	return res, nil
}

func (ins *LdapAdapter) Add(person *PersonModel) (err error) {
	req := ldap.AddRequest{
		DN: getDN(person.AccountName),
		Attributes: []ldap.Attribute{
			{Type: "objectClass", Vals: []string{"top", "inetOrgPerson"}},
			{Type: "uid", Vals: []string{person.AccountName}},
			{Type: "userPassword", Vals: []string{person.Password}},
			{Type: "cn", Vals: []string{person.FamilyName + person.GivenName}},
			{Type: "sn", Vals: []string{person.FamilyName}},
			{Type: "givenName", Vals: []string{person.GivenName}},
			{Type: "mail", Vals: []string{person.Mail}},
			{Type: "displayName", Vals: []string{person.DisplayName}},
		}}

	err = ins.conn.Add(&req)
	if err != nil {
		return NewError(err)
	}

	return nil
}

func (ins *LdapAdapter) Delete(accountName string) (err error) {
	// 搜索DN
	res, err := ins.searchOUAccountDN(accountName)
	if err != nil {
		return err
	}

	if res == nil {
		return NewError(ldap.NewError(ldap.LDAPResultNoSuchObject, nil))
	}

	err = ins.conn.Del(&ldap.DelRequest{DN: res.Entries[0].DN})
	if err != nil {
		return NewError(err)
	}

	return nil
}

func (ins *LdapAdapter) Find(accountName string) (bool, error) {

	res, err := ins.searchOUAccountDN(accountName)
	if err != nil {
		return false, NewError(err)
	}

	return res != nil, nil
}

func (ins *LdapAdapter) Get(accountName string) (*PersonModel, error) {
	res, err := ins.conn.Search(ldap.NewSearchRequest(
		ouBaseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=inetOrgPerson)(uid="+accountName+"))",
		[]string{"uid", "cn", "sn", "givenName", "mail", "displayName"}, // A list attributes to retrieve
		nil,
	))
	if err != nil {
		return nil, NewError(err)
	}

	if len(res.Entries) == 0 {
		return nil, NewError(ldap.NewError(ldap.LDAPResultNoSuchObject, err))
	} else if len(res.Entries) > 1 {
		return nil, NewErrorWithMessage(ldap.LDAPResultOther, "存在多个相同账户名")
	}

	entry := res.Entries[0]

	return &PersonModel{
		AccountName: entry.GetAttributeValue("uid"),
		DisplayName: entry.GetAttributeValue("displayName"),
		FamilyName:  entry.GetAttributeValue("sn"),
		GivenName:   entry.GetAttributeValue("givenName"),
		Mail:        entry.GetAttributeValue("mail"),
	}, nil
}

func (ins *LdapAdapter) Authentication(accountName string, accountPassword string) (err error) {

	// 搜索DN
	res, err := ins.searchOUAccountDN(accountName)
	if err != nil {
		return err
	}

	if res == nil {
		return NewError(ldap.NewError(ldap.LDAPResultNoSuchObject, nil))
	}

	// 用新的链接来认证，不影响原链接的绑定结果。
	newConn, err := ldap.DialURL(ldapAdderDefault)
	if err != nil {
		return NewError(err)
	}

	defer newConn.Close()

	userdn := res.Entries[0].DN

	err = newConn.Bind(userdn, accountPassword)
	if err != nil {
		return NewError(err)
	}

	return nil
}

// 重置密码
func (ins *LdapAdapter) ResetPassword(accountName string, accountPassword string) (err error) {

	err = ins.conn.Modify(&ldap.ModifyRequest{
		DN: getDN(accountName),
		Changes: []ldap.Change{
			{Operation: ldap.ReplaceAttribute,
				Modification: ldap.PartialAttribute{Type: "userPassword", Vals: []string{accountPassword}},
			},
		},
	})

	if err != nil {
		return NewError(err)
	}

	return nil
}
