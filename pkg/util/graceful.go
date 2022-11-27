package util

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GracefulStop 优雅停止
func GracefulStop(server http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, channel := context.WithTimeout(context.Background(), 10*time.Second)
	defer channel()
	go func() {
		time.Sleep(10 * time.Second)
		panic("UnGracefulStop")
	}()
	return server.Shutdown(ctx)
}
