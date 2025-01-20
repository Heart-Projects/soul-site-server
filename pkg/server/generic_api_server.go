package server

import (
	"com.sj/admin/middleware"
	"com.sj/admin/pkg/options"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type GenericAPIServer struct {
	// 安装的中间件
	middlewares []string
	// ShutdownTimeout 优雅关闭超时时间
	ShutdownTimeout time.Duration
	// Engine gin引擎
	*gin.Engine
	// healthz 是否启用健康检查
	healthz bool
	// enableMetrics 是否启用metrics
	enableMetrics bool
	//
	enableProfiling bool
	port            int
	instance        *http.Server
}

// New 创建一个 server
func New(option *options.ServerOptions) *GenericAPIServer {
	server := &GenericAPIServer{
		Engine:          gin.Default(),
		middlewares:     option.Middlewares,
		ShutdownTimeout: option.ShutdownTimeout,
		healthz:         option.Healthz,
		enableMetrics:   option.EnableMetrics,
		enableProfiling: option.EnableProfiling,
		port:            option.Port,
	}
	server.init()
	return server
}

// Start 启动 server
func (s *GenericAPIServer) init() {
	s.setup()
	s.setupMiddlewares()
	s.installApis()
	s.instance = &http.Server{
		Addr:           "localhost:" + strconv.Itoa(s.port),
		Handler:        s,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

}

func (s *GenericAPIServer) setup() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		logrus.Info("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

func (s *GenericAPIServer) setupMiddlewares() {
	s.Use(middleware.Cors())
	s.Use(middleware.Auth(s.Engine))
}

func (s *GenericAPIServer) installApis() {
	if s.healthz {
		s.GET("/healthz", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}
	if s.enableProfiling {
		profilingRouter := &(s.RouterGroup)
		profilingRouter.Group("/debug/pprof")
		{
			profilingRouter.GET("/", pprofHandler(pprof.Index))
			profilingRouter.GET("/cmdline", pprofHandler(pprof.Cmdline))
			profilingRouter.GET("/profile", pprofHandler(pprof.Profile))
			profilingRouter.POST("/symbol", pprofHandler(pprof.Symbol))
			profilingRouter.GET("/trace", pprofHandler(pprof.Trace))
			profilingRouter.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
			profilingRouter.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
			profilingRouter.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
			profilingRouter.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
			profilingRouter.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
			profilingRouter.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
		}
	}
	// 注册版本号
	s.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "1.0.0",
		})
	})
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Run 启动 server
func (s *GenericAPIServer) Run() {
	// 启动http server
	go func() {
		logrus.Info("server starting on port ", s.port)
		if err := s.instance.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号
	quitChan := make(chan os.Signal)
	// kill 默认会发送 syscall.SIGTERM
	// kill -2 发送 syscall.SIGINT, 等同于crl+c
	// kill -9 发送 syscall.SIGKILL, 强制杀死进程, 但是不能被捕获
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 通知到quitChan
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan
	logrus.Info("server preparing to shutdown")
	// 创建一个 5 秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
	defer cancel()

	if err := s.instance.Shutdown(ctx); err != nil {
		logrus.Fatal("server forced to shutdown: ", err)
	}
	logrus.Info("server exiting")
}

// Close 关闭 server
func (s *GenericAPIServer) Close() error {
	return nil
}

// Ping ping
func (s *GenericAPIServer) Ping() error {
	return nil
}
