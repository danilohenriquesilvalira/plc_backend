package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"plc_project/config"
	"plc_project/internal/cache"
	"plc_project/internal/database"
	"plc_project/internal/plcmanager"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewDB(cfg.MariaDB)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer db.Close()

	redis, err := cache.NewRedisCache(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password)
	if err != nil {
		log.Fatalf("Erro ao conectar ao Redis: %v", err)
	}
	defer redis.Close()

	logger := database.NewLogger(db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Println("Sinal de shutdown recebido, encerrando...")
		logger.Info("Shutdown", "Sinal de shutdown recebido")
		cancel()
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		plcmanager.RunAllPLCs(ctx, db, redis, logger)
	}()
	wg.Wait()

	log.Println("Aplicação encerrada.")
	logger.Info("Aplicação encerrada", "")
}
