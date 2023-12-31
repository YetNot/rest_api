package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	author2 "rest_api/internal/author"
	author "rest_api/internal/author/db"
	"rest_api/internal/config"
	"rest_api/internal/user"
	"rest_api/pkg/client/postgresql"
	"rest_api/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	postrgeSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	repository := author.NewRepository(postrgeSQLClient, logger)

	newAth := author2.Author{
		Name: "Mir",
	}
	if err = repository.Create(context.TODO(), &newAth); err != nil {
		logger.Fatalf("%v", err)
	}
	logger.Infof("%v", newAth)

	all, err := repository.FindAll(context.TODO())
	if err != nil {
		logger.Fatalf("%v", err)
	}

	for _, ath := range all {
		logger.Infof("%v", ath)
	}

	// cfgMongo := cfg.MongoDB
	// mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username,
	// 	cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	// if err != nil {
	// 	panic(err)
	// }
	// storage := db.NewStorage(mongoDBClient, cfg.MongoDB.Collection, logger)

	// user1 := user.User{
	// 	ID:           "",
	// 	Email:        "glinskix24@bk.ru",
	// 	Username:     "yetnot",
	// 	PasswordHash: "12345",
	// }
	// user1ID, err := storage.Create(context.Background(), user1)
	// if err != nil {
	// 	panic(err)
	// }
	// logger.Info(user1ID)

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket: %s", socketPath)

	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
