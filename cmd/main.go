package main

import (
	"fmt"
	stdlog "log"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ong-gtp/go-chat/pkg/config"
	"github.com/ong-gtp/go-chat/pkg/domain/middlewares"
	"github.com/ong-gtp/go-chat/pkg/models"
	"github.com/ong-gtp/go-chat/pkg/rabbitmq"
	"github.com/ong-gtp/go-chat/pkg/routes"
)

func main() {
	// Load env values
	err := godotenv.Load()
	if err != nil {
		stdlog.Fatal("Error loading .env file")
	}

	// Logging setup
	var logger log.Logger
	fileLogging := os.Getenv("LOG_TO_FILE")
	if fileLogging == "true" {
		file, err := os.Create(fmt.Sprintf("./applog-%s.txt", time.Now().Format(time.RFC3339Nano)))
		if err != nil {
			stdlog.Fatal("Could not create log file: ", err)
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
		stdlog.Fatal("JWT Secret not set")
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
	stdlog.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}
