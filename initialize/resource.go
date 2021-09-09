package initialize

import (
	"context"
	"strings"
	"wing/models/ent"
	"wing/models/ent/privacy"
	"wing/models/ent/resource"
	"wing/models/schema"
	"wing/models/util"
	"wing/resolver/resolvers"

	casbinadapter "wing/adapters/casbin"

	"github.com/golobby/container/v3"
)

type ResourceSet map[string]string

// NOTE 所有受限资源的资源数据描述
var resourceSet = ResourceSet{

	util.TToS(schema.System.Type):          "系统",
	util.TToS(schema.User.Type):            "用户",
	util.TToS(schema.Organization.Type):    "企业组织",
	util.TToS(schema.OrgUnit.Type):         "企业组织单元",
	util.TToS(schema.OrgUnitPosition.Type): "企业组织职务",
	util.TToS(schema.OrgUnitMember.Type):   "企业组织成员",
	util.TToS(schema.JobHistory.Type):      "就职履历",
	util.TToS(schema.Auth.Type):            "认证历史记录",
}

// 启动时更新权限系统使用的资源列表
func initResource(client *ent.Client) {

	// 忽略所有隐私策略
	ctxAllow := privacy.DecisionContext(context.Background(), privacy.Allow)

	tx, err := client.Tx(ctxAllow)
	if err != nil {
		panic(err.Error())
	}

	clientTx := tx.Client()

	err = migrateResourceAndPermission(ctxAllow, clientTx)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			panic(err)
		}
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}
}

// 取差分结果。
// 1、当前对象比other额外多出来的。
// 2、键相同，但值不同的。返回当前对象键值
func (ins ResourceSet) Difference(other ResourceSet) (extra ResourceSet, diff ResourceSet) {
	extra = make(ResourceSet)
	diff = make(ResourceSet)

	for k, v := range ins {
		ov, ok := other[k]

		// 额外的
		if !ok {
			extra[k] = v
		} else if strings.Compare(v, ov) != 0 {
			diff[k] = v
		}
	}
	return extra, diff
}

func addResourceAndPermission(ctx context.Context, casbin *casbinadapter.CasbinAdapter, client *ent.Client, extra ResourceSet) error {

	bulk := make([]*ent.ResourceCreate, len(extra))
	index := 0

	for typ, name := range extra {

		bulk[index] = client.Resource.Create().
			SetName(name).SetType(typ)

		index++
	}

	ress, err := client.Resource.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return err
	}

	orgId, err := client.Organization.Query().FirstID(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return err
	}

	posts, err := client.OrgUnitPosition.Query().All(ctx)
	if err != nil {
		return err
	}

	if len(posts) == 0 {
		return nil
	}

	// 对所有资源默人的所有动作都阻止
	for _, post := range posts {
		for _, res := range ress {
			for _, act := range casbinadapter.ActionAll {
				casbin.AddPolicyForRoleInDomain(resolvers.OrgDomainPrefix,
					orgId, post.ID, res.ID, act, casbinadapter.EffectDeny)
			}
		}
	}

	return nil
}

func deleteResourceAndPermission(ctx context.Context, casbin *casbinadapter.CasbinAdapter, client *ent.Client, extra ResourceSet) error {
	typs := make([]string, len(extra))
	idx := 0
	for k := range extra {
		typs[idx] = k
		idx++
	}

	ress := client.Resource.Query().Where(resource.TypeIn(typs...)).AllX(ctx)

	_, err := client.Resource.Delete().
		Where(resource.TypeIn(typs...)).
		Exec(ctx)
	if err != nil {
		return err
	}

	orgId, err := client.Organization.Query().FirstID(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return err
	}

	for _, res := range ress {
		casbin.DeletePoliciesForResourceInDomain(resolvers.OrgDomainPrefix,
			orgId, res.ID)
	}

	return nil
}

func updateResources(ctx context.Context, client *ent.Client, diff ResourceSet) error {

	for typ, name := range diff {
		err := client.Resource.Update().Where(resource.TypeEQ(typ)).SetName(name).Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// 检查资源变化情况，更新资源列表，并处理权限策略信息
func migrateResourceAndPermission(ctx context.Context, client *ent.Client) error {

	// 分别取代码模型和数据库中的资源信息

	resources, err := client.Resource.Query().All(ctx)
	if err != nil {
		return err
	}

	dbSet := make(ResourceSet)
	for _, resource := range resources {
		dbSet[resource.Type] = resource.Name
	}

	// 取差集，新资源插入数据库，更新名称变更的
	extra, diff := resourceSet.Difference(dbSet)

	var casbin *casbinadapter.CasbinAdapter
	err = container.Resolve(&casbin)
	if err != nil {
		return err
	}

	if len(diff) != 0 {
		err = updateResources(ctx, client, diff)
		if err != nil {
			return err
		}
	}

	if len(extra) != 0 {
		err = addResourceAndPermission(ctx, casbin, client, extra)
		if err != nil {
			return err
		}
	}

	// 旧资源删除，并删掉权限管理中的策略信息
	extra, _ = dbSet.Difference(resourceSet)
	if len(extra) != 0 {
		err := deleteResourceAndPermission(ctx, casbin, client, extra)
		if err != nil {
			return err
		}
	}

	return nil
}
