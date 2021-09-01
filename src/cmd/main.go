package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	_graphQLArticleDelivery "github.com/samudra-ajri/go-graphql/src/article/delivery/graphql"
	_articleRepo "github.com/samudra-ajri/go-graphql/src/article/repository/mysql"
	_articleUcase "github.com/samudra-ajri/go-graphql/src/article/usecase"
	_authorRepo "github.com/samudra-ajri/go-graphql/src/author/repository/mysql"
	"github.com/samudra-ajri/go-graphql/src/config"
	"github.com/samudra-ajri/go-graphql/src/middleware"
)

func init() {
	appName := config.GetConfig().AppName
	if appName == "local" || appName == "development" {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := config.GetConfig().DbHost
	dbPort := config.GetConfig().DbPort
	dbUser := config.GetConfig().DbUser
	dbPass := config.GetConfig().DbPassword
	dbName := config.GetConfig().DbName
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)
	authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	ar := _articleRepo.NewMysqlArticleRepository(dbConn)

	timeoutContext := time.Duration(2) * time.Second
	au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)

	schema := _graphQLArticleDelivery.NewSchema(_graphQLArticleDelivery.NewResolver(au))
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    schema.Query(),
		Mutation: schema.Mutation(),
	})
	if err != nil {
		logrus.Fatal(err)
	}

	graphQLHandler := handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		GraphiQL: true,
		Pretty:   true,
	})

	e.GET("/graphql", echo.WrapHandler(graphQLHandler))
	e.POST("/graphql", echo.WrapHandler(graphQLHandler))

	log.Fatal(e.Start(":" + config.GetConfig().AppPort))
}
