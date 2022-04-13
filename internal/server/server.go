package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/avanibbles/flowflow/internal/util"

	"github.com/avanibbles/flowflow/pkg/config"

	"github.com/avanibbles/flowflow/pkg/docs"

	"github.com/avanibbles/flowflow/internal/routes"
	"github.com/avanibbles/flowflow/internal/services"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/avanibbles/flowflow/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	config    *config.Config
	logger    *zap.Logger
	topLogger *zap.Logger
	wg        *sync.WaitGroup
	df        services.DependencyFactory
}

func (s *Server) toDependencyConfig() *services.DependencyConfig {
	return &services.DependencyConfig{
		Logger: s.topLogger,
		Config: s.config,
		Wg:     s.wg,
	}
}

func (s *Server) buildHttp() (*http.Server, error) {
	e := echo.New()
	e.Use(ZapLogger(s.logger, []string{"/health/*"}))
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = util.MakeEchoErrorHandler(s.logger.With(zap.String("component", "HttpErrorHandler")))

	docs.SwaggerInfo.Version = internal.Version
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	if err := routes.Setup(s.df, e); err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%s:%d", s.config.Http.Host, s.config.Http.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	return srv, nil
}

func (s *Server) getListenerFunc() (func(server *http.Server) error, error) {
	return func(server *http.Server) error {
		s.logger.Info("server listening", zap.String("addr", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	}, nil
}

func NewServer(config *config.Config) (*Server, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	logger := internal.GetLogger()
	wg := &sync.WaitGroup{}

	depCfg := &services.DependencyConfig{
		Logger: logger,
		Config: config,
		Wg:     wg,
	}
	deps, err := services.NewDependencyFactory(depCfg)
	if err != nil {
		return nil, err
	}

	return &Server{
		config:    config,
		topLogger: logger,
		logger:    logger.With(zap.String("component", "Server")),
		wg:        wg,
		df:        deps,
	}, nil
}

func (s *Server) Run() error {
	srv, err := s.buildHttp()
	if err != nil {
		return err
	}

	listener, err := s.getListenerFunc()
	if err != nil {
		return err
	}

	maintenance := s.df.GetDomain().NewMaintenanceService()

	if err := maintenance.PreServiceStart(); err != nil {
		return err
	}

	errChan := make(chan error, 1)
	go func() {
		if err := listener(srv); !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	if err := maintenance.OnServiceStart(); err != nil {
		errChan <- err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
		s.logger.Info("shutdown requested")
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			s.logger.Error("shutdown error", zap.Error(err))
			return err
		}

		s.wg.Wait()
		return nil

	case e := <-errChan:
		s.logger.Error("server error", zap.Error(err))
		return e
	}
}
