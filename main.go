package main

import (
	"errors"
	"fmt"
	stdlog "log"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ong-gtp/go-chat/config"
	"github.com/ong-gtp/go-chat/http/middlewares"
	"github.com/ong-gtp/go-chat/http/routes"
	"github.com/ong-gtp/go-chat/models"
	"github.com/ong-gtp/go-chat/services/rabbitmq"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() error {
	// Load env values
	err := godotenv.Load()
	if err != nil {
		stdlog.Println("Error loading .env file")
		return err
	}

	// Logging setup
	var logger log.Logger
	fileLogging := os.Getenv("LOG_TO_FILE")
	if fileLogging == "true" {
		file, err := os.Create(fmt.Sprintf("./applog-%s.txt", time.Now().Format(time.RFC3339Nano)))
		if err != nil {
			stdlog.Println("Could not create log file: ", err)
			return err
		}
		defer file.Close()

		// Logfmt is a structured, key=val logging format that is easy to read and parse
		logger = log.NewLogfmtLogger(log.NewSyncWriter(file))
	} else {
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	}
	// Direct any attempts to use Go's log package to our structured logger
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	// Log the timestamp (in UTC) and the callsite (file + line number) of the logging
	// call for debugging in the future.
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)

	// Setup Database and migrate models
	config.ConnectDB()
	db := config.GetDB()
	db.AutoMigrate(models.Tables...)
	level.Info(logger).Log("Database", "migrated")

	// Connect Rabbit MQ
	conn, ch := rabbitmq.InitilizeBroker(logger)
	defer conn.Close()
	defer ch.Close()

	// JWT_SECRET must be set for Auth signing
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		stdlog.Println("JWT Secret not set")
		return errors.New("JWT Secret not set")
	}

	// Setup app routes
	r := mux.NewRouter()
	routes.RegisterAuthRoutes(r)
	routes.RegisterChatRoutes(r)
	routes.RegisterWebsocketRoute(r)

	// Wrap routes with logging and cors middlewares
	loggingMiddleware := middlewares.LoggingMiddleware(logger)
	loggedRoutes := loggingMiddleware(r)
	handler := middlewares.Cors(loggedRoutes)

	// Start api server
	port := os.Getenv("PORT")
	level.Info(logger).Log("Server", "starting", "port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	return err
}
