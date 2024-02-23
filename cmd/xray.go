package cmd

import (
	"Txray/core/manage"
	"Txray/xray"
	"github.com/abiosoft/ishell"
	"strconv"
)

func InitServiceShell(shell *ishell.Shell) {
	// 启动或重启服务
	shell.AddCmd(&ishell.Cmd{
		Name: "run",
		Help: "启动或重启服务",
		Func: func(c *ishell.Context) {
			if len(c.Args) == 1 {
				xray.Start(c.Args[0])
			} else {
				xray.Start(strconv.Itoa(manage.Manager.SelectedIndex()))
			}

		},
	})
	// 停止服务
	shell.AddCmd(&ishell.Cmd{
		Name: "stop",
		Help: "停止服务",
		Func: func(c *ishell.Context) {
			xray.Stop()
		},
	})
	// Generate
	shell.AddCmd(&ishell.Cmd{
		Name: "gen",
		Help: "生成配置文件",
		Func: func(c *ishell.Context) {
			if len(c.Args) == 1 {
				xray.Generate(c.Args[0])
			} else {
				xray.Generate(strconv.Itoa(manage.Manager.SelectedIndex()))
			}

		},
	})
}
