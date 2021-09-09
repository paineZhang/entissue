#!/bin/bash

go run -mod=mod ./tools/entc.go

# NOTE 先不用下面的兼容性了。缺少“模板”的参数输入，官方文档写的不明白，也没找到参数的输入内容。
# 为了版本兼容性，不使用entc（ent的cli），使用代码中依赖的版本的实现。
# go run entgo.io/ent/cmd/ent generate --feature privacy,entql,schema/snapshot,sql/schemaconfig,sql/lock,sql/modifier,sql/upsert --target ./models/ent ./models/schema