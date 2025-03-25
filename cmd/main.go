package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"trood-test/api"
	"trood-test/clients/openai"
	"trood-test/db/postgres"
	"trood-test/env"
	"trood-test/internal/repository"
	"trood-test/internal/services/nlp"
	kafkaproducer "trood-test/kafka/producer"
)


const(
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func main(){
    env := env.MustLoad()
    log := setupLogger(env.Env)

    if env.Env == envLocal {
        log.Info("starting application", slog.Any("cfg", env))
    }

	pg, err := postgres.New(&env.PgSql)
	if err != nil {
        log.Error("Failed to connect to postgres", slog.String("error", err.Error()))
        os.Exit(1)
    }
	defer pg.Stop()

    repo := repository.New(pg)

    eventDispatcher, err := kafkaproducer.NewKafkaProducer(env.Env, env.Kafka, log)
    if err != nil {
        log.Error("Failed to create kafka producer", slog.String("error", err.Error()))
        os.Exit(1)
    }
    defer eventDispatcher.Close()

	openaiClient := openai.NewClient(&env.OpenaiClient)

	nlpService := nlp.New(log, repo, openaiClient, openaiClient)

	app := api.New(log, env, nlpService)

	quit := make(chan os.Signal, 1) 
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func ()  {
		if err := app.Run(); err != nil{
			log.Error("server error", slog.String("error", err.Error()))
			quit <- syscall.SIGTERM
		}
	}()
	
	log.Info("Application started successfully")

    <-quit
    log.Info("Shutdown signal received")
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := app.GracefulShutdown(ctx); err != nil {
        log.Error("Server shutdown error", slog.String("error", err.Error()))
    }
    
    log.Info("Application has been shutdown gracefully")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	log = slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	return log
}
