package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	port = flag.String("port", "8000", "the port to run server on")
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(logger.Writer())

	router := gin.Default()

	// just for test
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
		})
	})

	svr := &http.Server{
		Addr:           ":" + *port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go func() {
		logger.Printf("Starting server at http://localhost:%s", *port)

		err := svr.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			logger.Println("listen:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Println("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown:", err)
	}

	logger.Println("Server exiting")
}
