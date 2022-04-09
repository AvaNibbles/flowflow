package util

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func LoopUntilCancel(wg *sync.WaitGroup, logger *zap.Logger, intervalMs int, f func() error) {
	wg.Add(1)
	defer wg.Done()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	for {
		if err := f(); err != nil {
			logger.Error("unhandled worker error", zap.Error(err))
		}

		select {
		case <-quit:
			return
		case <-time.After(time.Duration(intervalMs) * time.Millisecond):
			continue
		}
	}
}
