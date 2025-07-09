package main

import (
	"context"
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

	srv := httptransport.New(httptransport.Options{
		Port:   cfg.Port,
		Logger: logg,
		Routes: nil,
	})

	if err := srv.Start(context.Background()); err != nil {
		logg.Fatal("server stopped with error", err)
	}
}
