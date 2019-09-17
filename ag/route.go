package ag

import (
	"strings"

	"github.com/bingoohuang/gou/str"

	"github.com/bingoohuang/gonet"
	"github.com/bingoohuang/gou/htt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (a *App) setupRoutes() {
	r := a.Gin

	r.GET("/*arg", a.Query)
	r.POST("/*arg", a.Exec)

	logrus.Infof("start to run at address %s", a.ListenAddr)
	if err := r.Run(a.ListenAddr); err != nil {
		logrus.Warnf("fail to start at %s, error %v", a.ListenAddr, err)
	}
}

func (a *App) Query(c *gin.Context) {
	key, query := a.parseArg(c)
	p := a.Processor(key)
	ok, err := p.Query(query)

	if err == nil {
		c.Data(200, htt.ContentTypeJSON, []byte(ok))
	} else {
		c.Data(500, htt.ContentTypeText, []byte(err.Error()))
	}
}

func (a *App) Exec(c *gin.Context) {
	key, query := a.parseArg(c)

	p := a.Processor(key)
	body := gonet.ReadString(c.Request.Body)
	ok, err := p.Exec(query, body)

	if err == nil {
		c.Data(200, htt.ContentTypeJSON, []byte(ok))
	} else {
		c.Data(500, htt.ContentTypeText, []byte(err.Error()))
	}
}

func (a *App) parseArg(c *gin.Context) (string, string) {
	arg := strings.TrimPrefix(c.Param("arg"), "/")
	pos := strings.Index(arg, "/")
	key := str.If(pos > 0, Before(arg, pos), arg)
	query := str.If(pos > 0, After(arg, pos+1), "")
	return key, query
}

func Before(s string, pos int) string {
	if pos >= 0 && pos < len(s) {
		return s[0:pos]
	}

	return ""
}
func After(s string, pos int) string {
	if pos >= 0 && pos < len(s) {
		return s[pos:]
	}
	return ""
}
