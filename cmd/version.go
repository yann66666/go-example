// Author: yann
// Date: 2021/12/22
// Desc: cmd

package cmd

import (
	"fmt"
	"github.com/Hidata-xyz/go-example/build"
	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{
	Name:  "version",
	Usage: "print version",
	Action: func(c *cli.Context) error {
		fmt.Println(build.Version())
		return nil
	},
}
