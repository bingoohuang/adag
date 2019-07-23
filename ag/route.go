package ag

import (
	"github.com/bingoohuang/gonet"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (a *App) setupRoutes() {
	r := a.Gin

	r.GET("/:key", a.Query)
	r.POST("/:key", a.Exec)

	logrus.Infof("start to run at address %s", a.ListenAddr)
	if err := r.Run(a.ListenAddr); err != nil {
		logrus.Warnf("fail to start at %s, error %v", a.ListenAddr, err)
	}
}

func (a *App) Query(c *gin.Context) {
	key := c.Param("key")
	p := a.Processor(key)
	body := gonet.ReadString(c.Request.Body)
	ok, err := p.Query(body)

	if err == nil {
		c.Data(200, "application/json; charset=utf-8", []byte(ok))
	} else {
		c.Data(500, "text/plain; charset=utf-8", []byte(err.Error()))
	}
}

func (a *App) Exec(c *gin.Context) {
	key := c.Param("key")
	p := a.Processor(key)
	body := gonet.ReadString(c.Request.Body)
	ok, err := p.Exec(body)

	if err == nil {
		c.Data(200, "application/json; charset=utf-8", []byte(ok))
	} else {
		c.Data(500, "text/plain; charset=utf-8", []byte(err.Error()))
	}
}
