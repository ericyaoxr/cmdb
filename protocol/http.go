package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	hostAPI "github.com/ericyaoxr/cmdb/app/host/http"
	searchAPI "github.com/ericyaoxr/cmdb/app/resource/http"
	secretAPI "github.com/ericyaoxr/cmdb/app/secret/http"
	taskAPI "github.com/ericyaoxr/cmdb/app/task/http"
	"github.com/ericyaoxr/cmdb/conf"
)

// NewHTTPService 构建函数
func NewHTTPService() *HTTPService {
	r := httprouter.New()

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.Addr(),
		Handler:           cors.AllowAll().Handler(r),
	}
	return &HTTPService{
		r:      r,
		server: server,
		l:      zap.L().Named("API"),
		c:      conf.C(),
	}
}

// HTTPService http服务
type HTTPService struct {
	r      *httprouter.Router
	l      logger.Logger
	c      *conf.Config
	server *http.Server
}

// Start 启动服务
func (s *HTTPService) Start() error {
	// 装置子服务路由
	hostAPI.RegistAPI(s.r)
	secretAPI.RegistAPI(s.r)
	taskAPI.RegistAPI(s.r)
	searchAPI.RegistAPI(s.r)

	// 启动 HTTP服务
	s.l.Infof("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service is stopped")
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

// Stop 停止server
func (s *HTTPService) Stop() error {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}
	return nil
}
