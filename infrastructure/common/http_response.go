package common

import (
	errcode2 "github.com/go-eagle/eagle/infrastructure/common/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"

	httpstatus "github.com/go-eagle/eagle/infrastructure/transport/http/status"
	"github.com/go-eagle/eagle/infrastructure/utils"
)

// Response define a response struct
type Response struct {
	Retcode int         `json:"retcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewResponse return a response
func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	c.JSON(http.StatusOK, Response{
		Retcode: errcode2.Success.Code(),
		Message: errcode2.Success.Msg(),
		Data:    data,
	})
}

// Error return a error response
func (r *Response) Error(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusOK, Response{
			Retcode: errcode2.Success.Code(),
			Message: errcode2.Success.Msg(),
			Data:    gin.H{},
		})
		return
	}

	if v, ok := err.(*errcode2.Error); ok {
		response := Response{
			Retcode: v.Code(),
			Message: v.Msg(),
			Data:    gin.H{},
		}
		c.JSON(errcode2.ToHTTPStatusCode(v.Code()), response)
		return
	} else {
		// receive gRPC error
		if st, ok := status.FromError(err); ok {
			response := Response{
				Retcode: int(st.Code()),
				Message: st.Message(),
				Data:    gin.H{},
			}
			details := st.Details()
			if len(details) > 0 {
				response.Message += "|| details:"
				for _, v := range details {
					response.Message += "{ " + cast.ToString(v) + "} "
				}
				response.Message += "||"
			}
			// https://httpstatus.in/
			// https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go#L15
			// https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
			c.JSON(httpstatus.HTTPStatusFromCode(st.Code()), response)
			return
		}
	}
}

// RouteNotFound 未找到相关路由
func RouteNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "the route not found")
}

// healthCheckResponse 健康检查响应结构体
type healthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// HealthCheck will return OK if the underlying BoltDB is healthy.
// At least healthy enough for demoing purposes.
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, healthCheckResponse{Status: "UP", Hostname: utils.GetHostname()})
}
