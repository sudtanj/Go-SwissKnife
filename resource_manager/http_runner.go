package resource_manager

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sudtanj/Go-SwissKnife/env"
	"github.com/sudtanj/Go-SwissKnife/log"
	"go.uber.org/zap"
	zapAdapter "logur.dev/adapter/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunGin(logger *zapAdapter.Logger, serv *http.Server) (err error) {
	err = serv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Error("cannot run http with error message", log.ToMapInterface(zap.Error(err)))
		return err
	}
	return nil
}

func WatchGin(serv *http.Server, logger *zapAdapter.Logger, sigintServer chan os.Signal) error {
	signal.Notify(sigintServer, syscall.SIGINT, syscall.SIGTERM)
	<-sigintServer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		logger.Error("Http server shutdown failed! message ", log.ToMapInterface(zap.Error(err)))
		return err
	}
	return nil
}

func RunWithRoutine[T env.IConfig](e *gin.Engine, config T, logger *zapAdapter.Logger) {
	serv := &http.Server{
		Addr:    fmt.Sprintf("%+v:%+v", config.GetAddr(), config.GetPort()),
		Handler: e,
	}
	logger.Info("init gin completed!")

	go func() {
		logger.Info("Run Gin")
		err := RunGin(logger, serv)
		if err != nil {
			logger.Error("startGin", log.ToMapInterface(zap.Error(err)))
		}
	}()

	osSignal := make(chan os.Signal, 1)
	logger.Info("Run Watch Gin")
	err := WatchGin(serv, logger, osSignal)
	if err != nil {
		logger.Error("stopGin", log.ToMapInterface(zap.Error(err)))
	}
}
