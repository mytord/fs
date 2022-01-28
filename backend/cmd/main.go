package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	openapi "github.com/mytord/fs/backend/gen/opencliapi"
	"github.com/mytord/fs/backend/internal"
	"github.com/mytord/fs/backend/internal/repositories"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

func main() {
	initLogger()

	db := initDatabase()
	defer db.Close()

	profileRep := repositories.NewProfileRepository(db)

	privateApiService := internal.NewPrivateApiService(profileRep)
	privateApiController := openapi.NewPrivateApiController(
		privateApiService,
		openapi.WithPrivateApiErrorHandler(internal.ErrorHandler),
	)

	publicApiService := internal.NewPublicApiService(profileRep)
	publicApiController := openapi.NewPublicApiController(
		publicApiService,
		openapi.WithPublicApiErrorHandler(internal.ErrorHandler),
	)

	router := internal.NewRouter(publicApiController.Routes(), privateApiController.Routes())

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	ctx := context.Background()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("failed to run server: %s", err)
		}
	}()

	zap.S().Info("server started")

	<-ctx.Done()

	zap.S().Info("stopping server")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctxShutDown); err != nil {
		zap.S().Fatalf("shutdown failed: %s", err)
	}

	zap.S().Info("server stopped")
}

func initLogger() {
	logger, err := zap.NewProduction()
	defer logger.Sync()

	if err != nil {
		panic("failed to initialize logger")
	}

	zap.ReplaceGlobals(logger)
}

func initDatabase() *sql.DB {
	dbHost, ok := os.LookupEnv("MYSQL_HOST")

	if !ok || dbHost == "" {
		panic("invalid MYSQL_HOST setting")
	}

	dbUser, ok := os.LookupEnv("MYSQL_USER")

	if !ok || dbUser == "" {
		panic("invalid MYSQL_USER setting")
	}

	dbPassword, ok := os.LookupEnv("MYSQL_PASSWORD")

	if !ok || dbPassword == "" {
		panic("invalid MYSQL_PASSWORD setting")
	}

	dbName, ok := os.LookupEnv("MYSQL_DATABASE")

	if !ok || dbName == "" {
		panic("invalid MYSQL_DATABASE setting")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(fmt.Sprintf("failed to initialize database - %s: %s", dsn, err))
	}

	return db
}
