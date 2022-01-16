package main

import (
	"database/sql"
	"fmt"
	openapi "github.com/mytord/fs/backend/gen/opencliapi"
	"github.com/mytord/fs/backend/internal"
	"github.com/mytord/fs/backend/internal/repositories"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"net/http"
	"os"
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

	zap.S().Debug("backend server started")

	zap.S().Fatal(http.ListenAndServe(":8080", router))
}

func initLogger() {
	logger, err := zap.NewDevelopment()
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
