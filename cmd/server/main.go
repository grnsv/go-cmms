package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grnsv/go-cmms/internal/api/handler"
	"github.com/grnsv/go-cmms/internal/app"
	"github.com/grnsv/go-cmms/internal/config"
	"github.com/grnsv/go-cmms/internal/infrastructure"
	"github.com/grnsv/go-cmms/internal/infrastructure/postgres/repository"
)

func main() {
	// 1. Загрузить конфигурацию
	cfg := config.Load()
	log.Printf("Config loaded: server=%s:%d, db=%s", cfg.Server.Host, cfg.Server.Port, cfg.Database.URL)

	// 2. Инициализировать БД
	queries, db, err := infrastructure.NewDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Database connected successfully")

	// 3. Создать репозитории
	equipmentRepo := repository.NewEquipmentRepository(queries)
	equipmentClassRepo := repository.NewEquipmentClassRepository(queries)
	uow := repository.NewUnitOfWork(queries)

	// 4. Создать use cases
	listEquipmentUC := app.NewListEquipmentUseCase(equipmentRepo)
	getEquipmentByIDUC := app.NewGetEquipmentByIDUseCase(equipmentRepo)
	createEquipmentUC := app.NewCreateEquipmentUseCase(equipmentRepo)

	// 5. Создать handler
	h := handler.NewHandler(
		listEquipmentUC,
		getEquipmentByIDUC,
		createEquipmentUC,
	)

	// 6. Создать и запустить HTTP сервер
	server := createServer(cfg.Server.Address(), h)
	
	go func() {
		log.Printf("Starting server on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// 7. Graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped")
	
	// Cleanup
	_ = uow
	_ = equipmentClassRepo
}

// createServer создаёт и конфигурирует HTTP сервер
func createServer(addr string, h *handler.Handler) *http.Server {
	mux := http.NewServeMux()

	// Примеры маршрутов (в реальности будет сгенерировано ogen)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"ok"}`)
	})

	mux.HandleFunc("/api/v1/equipment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"count":0,"items":[]}`)
			return
		}
		if r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, `{"id":"placeholder"}`)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/api/v1/equipment/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"id":"placeholder"}`)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	return &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
