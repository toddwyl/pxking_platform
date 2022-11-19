package http_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-eagle/eagle/global"
	"github.com/go-eagle/eagle/infrastructure/common"
	"github.com/go-eagle/eagle/infrastructure/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//type AdminApiController struct {
//	medalController admin.IMedalController
//}
//
//func NewAdminApiController() *AdminApiController {
//	return &AdminApiController{
//		medalController: admin.NewMedalController(),
//	}
//}

// NewHTTPServer creates a HTTP server
func NewHTTPServer(g *gin.Engine, c *global.ServerConfig) *http.Server {

	router := RegisterRouter(g)

	srv := http.NewServer(
		http.WithAddress(c.Addr),
		http.WithReadTimeout(c.ReadTimeout),
		http.WithWriteTimeout(c.WriteTimeout),
	)

	srv.Handler = router
	// NOTE: register svc to http server

	return srv
}

// RegisterRouter loads the middlewares, routes, handlers.
func RegisterRouter(g *gin.Engine) *gin.Engine {

	//controller := NewAdminApiController()

	// HealthCheck 健康检查路由
	g.GET("/health", common.HealthCheck)
	// metrics router 可以在 prometheus 中进行监控
	// 通过 grafana 可视化查看 prometheus 的监控数据，使用插件6671查看
	g.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// api router
	//api := g.Group("/")
	//api.GET("/api/admin/medal/search", handleFn(controller.medalController.GetMedalList))
	//api.Use(middleware.Auth())
	//{
	//	//api.GET("/users/:id/following", user.FollowList)
	//	//api.GET("/users/:id/followers", user.FollowerList)
	//}
	return g
}
