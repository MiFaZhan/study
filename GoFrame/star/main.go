package main

import (
	_ "star/internal/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"star/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	// 全局设置i18n
	g.I18n().SetLanguage("zh-CN")

	err := connDb()
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}

func connDb() error {
	err := g.DB().PingMaster()
	if err != nil {
		return err
	}
	return nil
}

func aa() {

}
