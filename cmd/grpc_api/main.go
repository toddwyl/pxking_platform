package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-eagle/eagle/application/grpc_controller"
	"github.com/go-eagle/eagle/global"
	eagle "github.com/go-eagle/eagle/infrastructure/app"
	"github.com/go-eagle/eagle/infrastructure/logger"
	"github.com/go-eagle/eagle/infrastructure/middleware"
	"time"
)

func main() {
	context.Background()
	// 初始化配置
	// init config
	cfg := global.InitConfig("./config", "local", "config")
	// -------------- init resource -------------
	logger.Init(&cfg.Logger)
	// init db
	global.InitDB(&cfg.Mysql)
	// init redis
	//redis.Init()

	gin.SetMode(cfg.Mode)
	g := gin.New()
	// 使用中间件
	g.Use(gin.Recovery(),
		middleware.NoCache,
		middleware.Options,
		middleware.Secure,
		middleware.Logging(),
		middleware.RequestID(),
		middleware.Metrics(cfg.Name),
		middleware.Tracing(cfg.Name),
		middleware.Timeout(3*time.Second),
	)

	// start app
	app := eagle.New(
		eagle.WithName(cfg.Name),
		eagle.WithVersion(cfg.Version),
		eagle.WithLogger(logger.GetLogger()),
		eagle.WithServer(
			// init http server
			//http_controller.NewHTTPServer(g, &cfg.HTTP),
			// init grpc server
			grpc_controller.NewGRPCServer(&cfg.GRPC),
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
