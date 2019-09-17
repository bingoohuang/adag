package ag

import (
	"strings"
	"time"

	"github.com/bingoohuang/cmd"
)

// BashProcessor bash执行处理器
type BashProcessor struct{}

func (s BashProcessor) Query(body string) (ok string, err error) {
	_, result := cmd.Bash(body, cmd.Timeout(5*time.Second))

	return strings.Join(result.Stdout, "\n"), result.Error
}

func (s BashProcessor) Exec(query string, body string) (ok string, err error) {
	return s.Query(body)
}
