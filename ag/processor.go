package ag

import (
	"fmt"
)

// Processor 表示抽象的处理器接口
type Processor interface {
	// Query 处理请求体body，成功返回ok, 失败返回err
	Query(body string) (ok string, err error)
	Exec(body string) (ok string, err error)
}

// RegisterProcess 注册处理器
func (a *App) RegisterProcessor(key string, processor Processor) error {
	if _, ok := a.Processors[key]; ok {
		return fmt.Errorf("key %s already registered", key)
	}

	a.Processors[key] = processor
	return nil
}

// Processor 获得处理器
func (a *App) Processor(key string) Processor {
	if p, ok := a.Processors[key]; ok {
		return p
	}

	return NotFoundProcessor{key: key}
}

// Processor 获得处理器
func (a *App) ExistsProcessor(key string) bool {
	_, ok := a.Processors[key]
	return ok
}

// NotFoundProcessor 没有发现处理器
type NotFoundProcessor struct{ key string }

func (n NotFoundProcessor) Query(_ string) (ok string, err error) {
	return "", fmt.Errorf("key %s not found", n.key)
}
func (n NotFoundProcessor) Exec(body string) (ok string, err error) {
	return n.Query(body)
}

// PingPongProcessor 乒乓处理器
type PingPongProcessor struct{}

func (p PingPongProcessor) Query(_ string) (ok string, err error) {
	return `{"data":"pong"}`, nil
}

func (p PingPongProcessor) Exec(body string) (ok string, err error) {
	return p.Query(body)
}
