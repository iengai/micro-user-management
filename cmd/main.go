package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"micro-user-management/cmd/config"
	"micro-user-management/cmd/interfaces/handler"
	router2 "micro-user-management/cmd/interfaces/router"
	"micro-user-management/infrastructure/db"
	"micro-user-management/infrastructure/persistence"
	"micro-user-management/internal/domain/usecase"
)

func main() {
	cfg := config.LoadConfig()
	mySQLDB, err := db.NewMySQLDB(cfg.DSN())
	if err != nil {
		panic(err)
	}
	defer func(mySQLDB *sql.DB) {
		err := mySQLDB.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}(mySQLDB)

	userRepo := persistence.NewUserRepository(mySQLDB)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	router := gin.Default()
	router2.SetUpRoutes(router, userHandler)

	go config.WatchConfig()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
