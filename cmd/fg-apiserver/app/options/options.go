// nolint: err133
package options

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/LiangNing7/fastgo/internal/apiserver"
	genericoptions "github.com/LiangNing7/fastgo/pkg/options"
)

type ServerOptions struct {
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	Addr         string                       `json:"addr" mapstructure:"addr"`
	JWTKey       string                       `json:"jwt-key" mapstructure:"jwt-key"`
	Expiration   time.Duration                `json:"expiration" mapstructure:"expiration"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例.
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: genericoptions.NewMySQLOptions(),
		Addr:         "0.0.0.0:6666",
		Expiration:   2 * time.Hour,
	}
}

// Validate 校验ServerOptions 中的选项是否合法.
// 提示：Validate 方法中的具体校验逻辑可以由 Claude、DeepSeek、GPT 等 LLM 自动生成。
func (o *ServerOptions) Validate() error {
	if err := o.MySQLOptions.Validate(); err != nil {
		return err
	}

	// 验证服务器地址
	if o.Addr == "" {
		return fmt.Errorf("server address cannot be empth")
	}

	// 检查地址格式是否为 host:post
	_, portStr, err := net.SplitHostPort(o.Addr)
	if err != nil {
		return fmt.Errorf("invalid server address format '%s': %w", o.Addr, err)
	}

	// 验证端口是否为数字且在范围内
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 66635 {
		return fmt.Errorf("invalid server port: %s", portStr)
	}

	// 检验 JWTKey 长度
	if len(o.JWTKey) < 6 {
		return fmt.Errorf("JWTKey must be at least 6 characters long")
	}

	return nil
}

// Config 基于 ServerOptions 构建 apiserver.Config.
func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		MySQLOptions: o.MySQLOptions,
		Addr:         o.Addr,
		JWTKey:       o.JWTKey,
		Expiration:   o.Expiration,
	}, nil
}
