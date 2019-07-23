package ag

import (
	"time"

	"github.com/bingoohuang/gou/str"

	"github.com/bingoohuang/now"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// App 表示Agent应用
type App struct {
	ListenAddr  string               // 监听端口，为空表示不启动监听。例如:9900
	AdServers   []string             // 服务端地址列表，例如[]string{"192.168.0.1:9901", "192.168.0.2:9901"}
	StartupTime now.Now              // 启动时间
	Gin         *gin.Engine          // Gin引擎
	Processors  map[string]Processor // 处理器
}

// CreateAgApp 创建AgApp应用。
// listenPort为监听端口，传0表示不启用监听。
// adServers 为admin服务器地址列表，例如192.168.0.1:9901,192.168.0.2:9901。不传递时表示本Agent不主动连接任何AdminServer
func CreateAgApp(listenAddr, adServers string) *App {
	agApp := &App{
		ListenAddr:  listenAddr,
		AdServers:   str.SplitTrim(adServers, ","),
		StartupTime: now.MakeNow(),
		Processors:  make(map[string]Processor),
	}

	if listenAddr != "" {
		gin.SetMode(gin.ReleaseMode)
		agApp.Gin = gin.Default()
	}

	return agApp
}

// GoStart 异步启动应用
func (a *App) GoStart() {
	if !a.ExistsProcessor("ping") {
		_ = a.RegisterProcessor("ping", &PingPongProcessor{})
	}

	if a.ListenAddr != "" {
		go a.setupRoutes()
	}

	if len(a.AdServers) > 0 {
		go a.setupAdServers()
	}
}

func (a *App) setupAdServers() {
	logrus.Infof("starting to scheduling to servers")
	d := 10 * time.Second
	timer := time.NewTimer(d)
	defer timer.Stop()

	for range timer.C {
		logrus.Infof("timer triggered")
		a.tapAdServers()
		timer.Reset(d)
	}
}

func (a *App) tapAdServers() {
	for _, adServer := range a.AdServers {
		a.tapAdServer(adServer)
	}
}

func (a *App) tapAdServer(adServer string) {

}
