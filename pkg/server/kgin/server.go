package kgin

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"github.com/guojiangli/picasso/pkg/klog/baselogger"

	"github.com/guojiangli/picasso/pkg/server"
	"github.com/guojiangli/picasso/pkg/utils/kstring"
	"github.com/gin-gonic/gin"
)

// Server ...
type Server struct {
	*gin.Engine
	Logger baselogger.Logger
	Server *http.Server
	mu     sync.Mutex
	info   server.ServiceInfo
}

func NewServer(opts ...*Option) (s *Server, err error) {
	ginopts := defaultOption().MergeOption(opts...)
	gin.SetMode(gin.ReleaseMode)
	s = &Server{
		Logger: ginopts.Logger,
	}
	if ginopts.Port == 0 {
		err = errors.New("Port不能为空")
		s.Logger.Log(err)
		return s, err
	}
	s.info = server.ServiceInfo{
		Name: "GIN",
		Host: ginopts.Host,
		Port: ginopts.Port,
	}
	s.Engine = gin.New()
	for _, m := range ginopts.Middlewares {
		s.Engine.Use(m)
	}
	return s, nil
}

// Serve implements server.Server interface.
func (s *Server) Serve() error {
	for _, route := range s.Engine.Routes() {
		s.Logger.Log("kgin_route:", kstring.KVString("method", route.Method), kstring.KVString("path", route.Path))
	}
	s.mu.Lock()
	s.Server = &http.Server{
		Addr:    s.info.Address(),
		Handler: s,
	}
	s.mu.Unlock()
	return s.Server.ListenAndServe()
}

// Stop implements server.Server interface
// it will terminate gin server immediately
func (s *Server) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Server.Close()
}

// GracefulStop implements server.Server interface
// it will stop gin server gracefully
func (s *Server) GracefulStop(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Server.Shutdown(ctx)
}

func (s *Server) Info() *server.ServiceInfo {
	return &s.info
}
