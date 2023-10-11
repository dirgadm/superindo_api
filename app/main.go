package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_productHandler "project-version3/superindo-task/product/delivery/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	// _articleHttpDelivery "github.com/bxcodec/go-clean-arch/article/delivery/http"
	// _articleHttpDeliveryMiddleware "github.com/bxcodec/go-clean-arch/article/delivery/http/middleware"
	_productRepo "project-version3/superindo-task/product/repository/mysql"
	_productUseCase "project-version3/superindo-task/product/usecase"
	// _articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	// _authorRepo "github.com/bxcodec/go-clean-arch/author/repository/mysql"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
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

	// Middleware Usage
	// middL := _articleHttpDeliveryMiddleware.InitMiddleware()
	// e.Use(middL.CORS)

	// authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	productRepo := _productRepo.NewMysqlProductRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	productUsecase := _productUseCase.NewProductUsecase(productRepo, timeoutContext)
	_productHandler.NewProductHandler(e, productUsecase)

	log.Fatal(e.Start(viper.GetString("server.address"))) //nolint
}
