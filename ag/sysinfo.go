package ag

import (
	"github.com/bingoohuang/gou/enc"
	"github.com/bingoohuang/sysinfo"
)

// SysinfoProcessor 系统信息处理器
type SysinfoProcessor struct{}

func (s SysinfoProcessor) Query(query string) (ok string, err error) {
	return enc.JSONPretty(sysinfo.GetSysInfo()), nil
}

func (s SysinfoProcessor) Exec(query string, body string) (ok string, err error) {
	return s.Query(body)
}
