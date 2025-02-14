package api

import (
	"log"
	"time"

	"plc_project/internal/cache"
	"plc_project/internal/database"
	"plc_project/internal/middleware"
	"plc_project/internal/websocket"

	"github.com/gorilla/mux"
)

// SetupRoutes configura todas as rotas da API.
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Adiciona middlewares globais
	router.Use(middleware.CORS())
	router.Use(middleware.RequestLogger())

	// Inicializa conexões necessárias
	db, _, err := getDBAndLogger()
	if err != nil {
		log.Fatalf("Erro ao inicializar banco: %v", err)
	}

	// Configura o Redis com retry
	var redis *cache.RedisCache
	maxRetries := 5
	retryDelay := 2 * time.Second
	for i := 0; i < maxRetries; i++ {
		redis, err = cache.NewRedisCache("localhost", 6379, "")
		if err == nil {
			log.Println("Conexão com Redis estabelecida com sucesso")
			break
		}
		log.Printf("Tentativa %d de %d de conectar ao Redis: %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(retryDelay)
		}
	}
	if err != nil {
		log.Fatalf("Erro ao inicializar Redis após %d tentativas: %v", maxRetries, err)
	}

	// Inicializa o gerenciador de WebSocket
	wsGerenciador := websocket.NovoGerenciador(db, redis)
	go wsGerenciador.Iniciar()

	// Configura as rotas
	setupAuthRoutes(router, db)
	setupPLCRoutes(router) // Removido o parâmetro db que não estava sendo usado

	// WebSocket endpoint
	router.HandleFunc("/ws/status", wsGerenciador.ManipularWS)

	return router
}

// setupAuthRoutes configura as rotas de autenticação.
func setupAuthRoutes(router *mux.Router, db *database.DB) {
	// Rota pública para login - apenas CORS e logging (já aplicados globalmente)
	router.HandleFunc("/api/auth/login",
		LoginHandler(db)).
		Methods("POST", "OPTIONS")

	// Rotas protegidas para gerenciamento de usuários
	router.HandleFunc("/api/auth/register",
		middleware.Auth(middleware.RequireRole("superadmin")(CreateUserHandler(db)))).
		Methods("POST", "OPTIONS")

	router.HandleFunc("/api/auth/update/{id}",
		middleware.Auth(middleware.RequireRole("superadmin")(UpdateUserHandler(db)))).
		Methods("PUT", "OPTIONS")

	router.HandleFunc("/api/auth/delete/{id}",
		middleware.Auth(middleware.RequireRole("superadmin")(DeleteUserHandler(db)))).
		Methods("DELETE", "OPTIONS")
}

// setupPLCRoutes configura as rotas de PLC e Tags.
func setupPLCRoutes(router *mux.Router) { // Removido o parâmetro db
	// Rotas de PLC
	router.HandleFunc("/api/plcs",
		middleware.Auth(GetPLCs)).
		Methods("GET", "OPTIONS")

	router.HandleFunc("/api/plcs/{id}",
		middleware.Auth(GetPLC)).
		Methods("GET", "OPTIONS")

	router.HandleFunc("/api/plcs",
		middleware.Auth(middleware.RequireRole("admin", "supervisor")(CreatePLC))).
		Methods("POST", "OPTIONS")

	router.HandleFunc("/api/plcs/{id}",
		middleware.Auth(middleware.RequireRole("admin", "supervisor")(UpdatePLC))).
		Methods("PUT", "OPTIONS")

	router.HandleFunc("/api/plcs/{id}",
		middleware.Auth(middleware.RequireRole("admin")(DeletePLC))).
		Methods("DELETE", "OPTIONS")

	// Rotas de Tags
	router.HandleFunc("/api/plcs/{id}/tags",
		middleware.Auth(GetTags)).
		Methods("GET", "OPTIONS")

	router.HandleFunc("/api/plcs/{id}/tags",
		middleware.Auth(middleware.RequireRole("admin", "supervisor")(CreateTag))).
		Methods("POST", "OPTIONS")

	router.HandleFunc("/api/tags/{tagId}",
		middleware.Auth(middleware.RequireRole("admin", "supervisor")(UpdateTag))).
		Methods("PUT", "OPTIONS")

	router.HandleFunc("/api/tags/{tagId}",
		middleware.Auth(middleware.RequireRole("admin")(DeleteTag))).
		Methods("DELETE", "OPTIONS")
}
