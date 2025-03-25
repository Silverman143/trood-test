package api

import (
	"context"
	"log/slog"
	"trood-test/api/server"
	"trood-test/env"
	"trood-test/internal/services/nlp"
)




type Api struct {
	log    *slog.Logger
	server *server.Server
	port string
}

func New(
	log *slog.Logger,
	env *env.Env,
	nlpService *nlp.NLPService) *Api {

	port := env.Http.GetPort()
	router := server.NewHandler(log, nlpService)
	routes := router.InitRouts()
	srv := server.NewServer(port, routes)

	return &Api{
		log:    log,
		server: srv,
		port: port,
	}
}

func (a *Api) Run() error {
	a.log.Info("Starting the application", slog.String("port", a.port))
	return a.server.Run()
}

func (a *Api) GracefulShutdown(ctx context.Context) error {
	a.log.Info("Initiating graceful shutdown")
	return a.server.Shutdown(ctx)
}