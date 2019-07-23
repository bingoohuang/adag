package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bingoohuang/adag/ag"
	"github.com/bingoohuang/adag/util"
	"github.com/bingoohuang/gou/http"
	"github.com/bingoohuang/gou/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	parseFlags()

	agListenAddr := viper.GetString("agListenAddr")
	adServers := viper.GetString("adServers")
	if agListenAddr != "" || adServers != "" {
		agApp := ag.CreateAgApp(agListenAddr, adServers)
		agApp.GoStart()
	}

	select {}
}

func parseFlags() {
	help := pflag.BoolP("help", "h", false, "help")
	pflag.StringP("agListenAddr", "", "", "agent listen address, eg :9900")
	pflag.StringP("adServers", "", "", "agent listen address, eg :9900")
	pprofAddr := http.PprofAddrPflag()
	pflag.Parse()
	args := pflag.Args()
	if len(args) > 0 {
		fmt.Printf("Unknown args %s\n", strings.Join(args, " "))
		pflag.PrintDefaults()
		os.Exit(0)
	}
	if *help {
		fmt.Printf("Built on %s from sha1 %s\n", util.Compile, util.Version)
		pflag.PrintDefaults()
		os.Exit(0)
	}
	http.StartPprof(*pprofAddr)

	// 绑定命令行参数，（优先级比配置文件高）
	log.Err(viper.BindPFlags(pflag.CommandLine))
}
