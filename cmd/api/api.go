package api

import (
	"direct/internal/config"
	list "direct/internal/service/client_list"
	stat "direct/internal/service/stat_client"
	"direct/pkg/logger"
	"direct/pkg/logger/middleware"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type ApiServer struct {
	addr string
	db   *sqlx.DB
}

func NewApiServer(addr string, db *sqlx.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {

	mux := http.NewServeMux()

	storeStat := stat.NewStore(s.db)
	statHandler := stat.NewHandler(storeStat)
	statHandler.RegisterRoutes(mux)

	storClientList := list.NewStore(s.db)
	listHandler := list.NewHandler(storClientList)
	listHandler.RegisterRoutes(mux)

	loggerJournal, err := logger.NewLogger(s.db)
	if err != nil {
		log.Fatalln("Failed to initialize logger:", err)
	}
	loggerMux := middleware.LoggerMiddleware(mux, loggerJournal)
	cors := config.Cors(loggerMux)

	if err := loggerJournal.LoggerBasic(logger.INFO_LOG, "Server started on port 8060"); err != nil {
		log.Println("Failed to log to database:", err)
	}

	fmt.Println("Server started on port 8060")

	if err := http.ListenAndServe("localhost:8060", cors); err != nil {
		loggerJournal.LoggerBasic(logger.ERROR_LOG, err.Error())
	}

	return err
}
