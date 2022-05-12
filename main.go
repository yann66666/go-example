// Author: yann
// Date: 2021/12/22
// Desc: XHiLand

package main

import (
	"fmt"
	"github.com/Hidata-xyz/go-example/build"
	"github.com/Hidata-xyz/go-example/cmd"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

const (
	service = "Hi-Sync"
)

func init() {

}

func main() {
	local := []*cli.Command{
		withCategory("command", cmd.RunCmd),
		withCategory("command", cmd.VersionCmd),
	}
	app := &cli.App{
		Name:                 service,
		Usage:                "synchronous data",
		Version:              build.Version(),
		EnableBashCompletion: true,
		Commands:             local,
	}
	app.Setup()
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("启动服务失败: %v", err)
		os.Exit(-1)
	}
}

func withCategory(cat string, cmd *cli.Command) *cli.Command {
	cmd.Category = strings.ToUpper(cat)
	return cmd
}
