# 编程语言

[golang](https://golang.google.cn)

# 框架/库

- [gin](https://gin-gonic.com)。服务框架。
- [gorm](https://gorm.io)。目前是间接的依赖，在 casbin 中，使用的适配器中依赖了 gorm
- [ent]()。图概念的数据库 ORM 库。其和 gqlgen 可以较好的集成（主要是定义的数据模型复用）。
- [gqlgen]()。graphql 库，主要应用在接口。

# 编辑器

VS code

# 代码生成

## 数据模型的生成

1. 使用 `sh ./tools/create_db_models.sh` 生成创建模型对象，并编辑模型内容
1. 使用 `sh ./tools/generate_db_models.sh` 生成模型对象的相关代码
1. 使用 `sh ./tools/check_db_models.sh` 通过打印信息检查模型正确性

## graph api 生成

1. 根据需要，修改`/tools/gqlgen.yml`
1. 创建/编辑 `gateway/graphql/api`目录下的 `*.graphql` 文件定义接口
1. 使用 `sh ./tools/generate_graphql_api.sh` 生成接口对象和处理方法

# 测试

## 单元测试

1. 增加/修改相关的`*_test.go`s 文件
1. 使用 `tools/unit_test.sh`进行所有模块单元测试的驱动，这里会排除不需要进行单元测试的测试文件(\*\_test.go)。

## api 测试

1. 使用 `tools/generate_api_test.sh`建立 api 测试使用的 client 和 query 代码。
1. 增加/修改 `api/test/graphql` 目录下的`*_test.go`文件。
1. 使用 `tools/api_test.sh`进行 api 的测试，这需要存在一个已经启动了的服务。

> api 测试使用的是`github.com/Yamashou/gqlgenc`，其依赖 gqlgen。
> gqlgenc 命令还不能指定配置文件的位置。所以，其配置文件`gqlgenc.yml`只能现放到根目录里。

# 其他

## 包依赖问题

- gqlparser 的版本暂时不要升到 2.2.0。这个版本增加了对 graphql 自定义指令的支持，但其生成的代码有问题，编译不过。
