package core

import (
	"net/http"

	"github.com/LiangNing7/fastgo/internal/pkg/errorsx"
	"github.com/gin-gonic/gin"
)

// ErrorResponse 定义了错误响应的结构，
// 用于 API 请求中发生错误时返回统一的格式化错误信息.
type ErrorResponse struct {
	// Reason 错误原因，标识错误类型.
	Reason string `json:"reason,omitempty"`

	// 错误详细的描述信息.
	Message string `json:"message,omitempty"`
}

func WriteResponse(c *gin.Context, data any, err error) {
	if err != nil {
		// 如果发生错误，生成错误响应.
		errx := errorsx.FromError(err)
		c.JSON(errx.Code, ErrorResponse{
			Reason:  errx.Reason,
			Message: errx.Message,
		})
		return
	}
	// 如果没有错误，返回成功响应.
	c.JSON(http.StatusOK, data)
}
