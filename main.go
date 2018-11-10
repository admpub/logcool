package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/admpub/logcool/cmd"
	_ "github.com/admpub/logcool/filter/grok"
	_ "github.com/admpub/logcool/filter/metrics"
	_ "github.com/admpub/logcool/filter/split"
	_ "github.com/admpub/logcool/filter/zeus"
	_ "github.com/admpub/logcool/input/collectd"
	_ "github.com/admpub/logcool/input/file"
	_ "github.com/admpub/logcool/input/http"
	_ "github.com/admpub/logcool/input/stdin"
	_ "github.com/admpub/logcool/output/email"
	_ "github.com/admpub/logcool/output/lexec"
	_ "github.com/admpub/logcool/output/redis"
	_ "github.com/admpub/logcool/output/stdout"
	"github.com/admpub/logcool/utils"
)

// config
var (
	conf    = flag.String("config", "", "path to config.json file")
	command = flag.String("command", "", "run in command, stdin2stdout.")
	custom  = flag.String("custom", "", "input custom template.")
	version = flag.Bool("version", false, "show version number.")
	std     = flag.Bool("std", false, "run in stadin/stdout.")
	help    = flag.Bool("help", false, "haha,I know you need me.")
)

func main() {
	flag.Parse()

	if *version != false {
		cmd.Version()
		os.Exit(0)
	}

	if *help != false {
		flag.Usage()
		os.Exit(0)
	}
	ctx, cancel := context.WithCancel(context.Background())
	var confs []utils.Config

	if *std != false {
		// cmd.Logcool()
		conf, err := utils.LoadDefaultConfig(ctx)
		if err != nil {
			fmt.Println(err)
		}
		confs = append(confs, conf)
	} else if *custom != "" {
		confs = cmd.Custom(ctx, *custom)
	} else if *command != "" {
		confs = cmd.Command(ctx, *command)
	} else {
		confs = cmd.LoadTemplates(ctx)
	}

	cmd.Run(confs)

	// 捕获ctrl-c,平滑退出
	chExit := make(chan os.Signal, 1)
	signal.Notify(chExit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	select {
	case <-chExit:
		cancel()
		fmt.Println("logcool EXIT...Bye.")
	}
}
