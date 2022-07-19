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
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(logger.Writer())

	if err := godotenv.Load(); err != nil {
		logrus.Fatal("could not get evn vars", err)
	}

	db, err := database.NewPostgresDB(database.Config{
		DBName:   os.Getenv("POSTGRES_DB"),
		Host:     "localhost",
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		SSLMode:  "disable",
		Username: os.Getenv("POSTGRES_USER"),
	})
	if err != nil {
		logger.Fatal("could not connect to the DB", err)
	}

	repos := provider.NewRepository(db)
	services := provider.NewService(repos)
	handlers := provider.NewHandler(services)

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
