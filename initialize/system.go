package initialize

import (
	"context"
	"wing/models/ent"
	"wing/models/ent/privacy"
)

func initSystemInfo(client *ent.Client) {
	// 忽略所有隐私策略
	ctxAllow := privacy.DecisionContext(context.Background(), privacy.Allow)

	count, err := client.System.Query().Count(ctxAllow)
	if err != nil {
		panic(err)
	}

	if count == 1 {
		return
	}

	if count > 1 {
		panic("存在多个系统信息。目前还不支持系统多租户")
	}

	_, err = client.System.Create().SetName("Wing").Save(ctxAllow)
	if err != nil {
		panic(err)
	}
}
