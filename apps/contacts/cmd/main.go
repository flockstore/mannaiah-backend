package main

import (
	"context"
	"github.com/flockstore/mannaiah-backend/apps/contacts/http"
	"github.com/flockstore/mannaiah-backend/apps/contacts/repository"
	"github.com/flockstore/mannaiah-backend/apps/contacts/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"log"

	appconfig "github.com/flockstore/mannaiah-backend/apps/contacts/config"
	commonconfig "github.com/flockstore/mannaiah-backend/common/config"
	"github.com/flockstore/mannaiah-backend/common/logger"
	httptransport "github.com/flockstore/mannaiah-backend/common/transport/http"
)

func main() {

	cfg, _, err := commonconfig.Load[appconfig.Config]("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logg := logger.New(cfg.LogLevel, nil)

	db, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	repo := repository.NewPostgresContactRepository(db)
	svc := service.NewContactService(repo)
	handler := http.New(svc)

	srv := httptransport.New(httptransport.Options{
		Port:   cfg.Port,
		Logger: logg,
		Routes: func(router fiber.Router) {
			handler.RegisterRoutes(router)
		},
	})

	if err := srv.Start(context.Background()); err != nil {
		logg.Fatal("server stopped with error", err)
	}

}
