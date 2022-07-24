package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-feedback/database"
	"product-feedback/provider"
	"product-feedback/server"
	"product-feedback/validation"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	port = flag.String("port", "8000", "the port to run server on")
)

func main() {
	flag.Parse()

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(logger.Writer())

	if err := godotenv.Load(); err != nil {
		logrus.Warn("could not get evn vars", err)
	}

	connStr := os.Getenv("DATABASE_URL")
	db, err := database.NewPostgresDB(connStr)
	if err != nil {
		logger.Fatal("could not connect to the DB", err)
	}

	v := validation.NewValidation()

	repos := provider.NewRepository(db)
	services := provider.NewService(repos)
	handlers := provider.NewHandler(logger, v, services)

	svr := server.NewServer(*port, handlers)

	go func() {
		logger.Printf("starting server at http://localhost:%s", *port)

		err := svr.Run()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			logger.Println("listen:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Println("shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		logger.Fatal("server forced to shutdown", err)
	}

	if err := db.Close(); err != nil {
		logger.Fatal("sould not close connection to the DB", err)
	}

	logger.Println("server exiting")
}
