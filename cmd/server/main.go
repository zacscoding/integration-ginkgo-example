package main

import (
	"flag"
	"fmt"
	"integration-ginkgo-example/internal/account"
	accountModel "integration-ginkgo-example/internal/account/model"
	"integration-ginkgo-example/internal/config"
	"integration-ginkgo-example/internal/database"
	"integration-ginkgo-example/pkg/logging"
	"net/http"
	"os"
	"time"
)

func main() {
	flagConfig := flag.String("config", "./config/config.yaml", "path to the config file")
	flag.Parse()
	logger := logging.DefaultLogger()

	// load application configurations
	conf, err := config.Load(*flagConfig)
	if err != nil {
		logger.Errorw("failed to load application configuration", "err", err)
		os.Exit(-1)
	}

	// setup database
	db, err := database.NewDatabase(conf.Database)
	if err != nil {
		logger.Fatalw("failed to load database", "err", err)
	}
	defer db.Close()
	if conf.Database.CreateTable {
		db.AutoMigrate(&accountModel.Account{})
	}

	// server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.Server.Port),
		Handler:      account.NewHandler(db),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Errorw("failed to start server", "err", err)
	}
}
