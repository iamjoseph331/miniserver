package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/iamjoseph331/miniserver/config"
	"github.com/iamjoseph331/miniserver/log"
	"github.com/iamjoseph331/miniserver/server/core"
	"github.com/iamjoseph331/miniserver/server/http"
	"github.com/iamjoseph331/miniserver/server/view"
	"golang.org/x/sync/errgroup"
)

var (
	httpServer *http.Server
)

func main() {
	initialize()

	wg, ctx := errgroup.WithContext(context.Background())
	wg.Go(httpServerStart)
	// signal processing
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
		log.Logger.Info(log.ApplicationLog(ctx, "received SIGTERM, exiting server gracefully..."))
	case <-ctx.Done():
	}
	// graceful shutdown
	gracefulShutdown()
}

func initialize() {
	config.Setup()
	log.Setup()

	// Initialize the core service
	coreService := core.serverCore()
	// Initialize the view
	view := view.serverCore(coreService)

	// Create new HTTP server
	httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%s", config.Conf.HTTPServer.Port),
		Handler:           serverhttp.NewHTTPServer(view),
		ReadHeaderTimeout: 20 * time.Second, // Set a limit on the time it takes to receive headers
	}

}

func httpServerStart() error {
	log.Logger.Info(log.ApplicationLog(context.Background(), "http server listening at :%s", config.Conf.HTTPServer.Port))
	return httpServer.ListenAndServe()
}

func gracefulShutdown() {
	ctx := context.Background()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Logger.Error(log.ApplicationLog(ctx, "Server forced to shutdown: %s", err.Error()))
	}
}
