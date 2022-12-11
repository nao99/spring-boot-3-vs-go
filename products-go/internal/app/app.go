package app

import (
	"fmt"
	"github.com/nao99/spring-boot-3-vs-go/products-go/internal/api"
	"github.com/nao99/spring-boot-3-vs-go/products-go/internal/config"
	"github.com/nao99/spring-boot-3-vs-go/products-go/internal/domain"
	"github.com/nao99/spring-boot-3-vs-go/products-go/internal/persistence/api"
	"github.com/nao99/spring-boot-3-vs-go/products-go/internal/server"
	"github.com/nao99/spring-boot-3-vs-go/products-go/pkg/db"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	srv := server.NewServer()

	dbSession := openDbSession(cfg)

	productRepository := repository.ProductRepositoryImpl{Session: dbSession}
	productService := domain.ProductServiceImpl{Repository: &productRepository}
	productController := api.ProductController{Service: &productService}

	api.SetupRoutes(srv, &productController)

	go func() {
		err = srv.Run(&cfg.Server)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	err = srv.Stop()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func openDbSession(cfg *config.Config) *gorm.DB {
	session, openingSessionError := db.OpenDbSession(cfg.Datasource.Dsn)
	if openingSessionError != nil {
		os.Exit(1)
	}

	return session
}
