// Author: yann
// Date: 2021/12/22
// Desc: cmd

package cmd

import (
	"context"
	"fmt"
	http_lib "github.com/HiData-xyz/HiData.More/HiExtern/lib/http"
	"github.com/HiData-xyz/HiData.More/HiExtern/lib/logging"
	"github.com/Hidata-xyz/go-example/internal"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var log = logging.Logger("run")

var RunCmd = &cli.Command{
	Name:  "run",
	Usage: "start manager",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "ip",
			Usage: "listen on address",
			Value: "0.0.0.0",
		},
		&cli.IntFlag{
			Name:  "port",
			Usage: "listen on address",
			Value: 80,
		},
		&cli.StringFlag{
			Name:    "gateway",
			Usage:   "所属宿主机gateway的ID",
			Value:   "uwp",
			EnvVars: []string{"GATEWAY_ID"},
		},
	},
	Action: func(c *cli.Context) error {
		address := fmt.Sprintf("%s:%d", c.String("ip"), c.Int("port"))
		//初始化http server
		httpServer := http_lib.Init(&http_lib.ServerConfig{
			Proto:    "http",
			Address:  address,
			CertPath: "",
			KeyPath:  "",
		})
		internal.InitRouter(httpServer.Engine)
		//启动http server
		var err error
		if err = httpServer.Start(context.Background()); err != nil {
			log.Errorf("server启动失败:%v", err)
			return err
		}
		log.Infof("服务启动: %v", address)
		Shutdown(func() {
			err = httpServer.Stop(context.Background())
		})
		return err
	},
}

//Shutdown 服务停止
func Shutdown(f func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	select {
	case s := <-c:
		log.Infof("收到停止信号: %v", s)
	}
	f()
	time.Sleep(time.Second * 3)
}
