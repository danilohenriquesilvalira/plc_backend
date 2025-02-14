// internal/api/auth.go
package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"plc_project/internal/database"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type contextKey string

const userContextKey contextKey = "user"

var jwtKey = []byte("sua_chave_secreta_aqui") // Em produção, use uma variável de ambiente

// Claims representa os dados que serão codificados no token JWT.
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// Estruturas de requisição e resposta:

// LoginRequest representa as credenciais enviadas para login.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse representa a resposta do login.
type LoginResponse struct {
	Token string        `json:"token"`
	User  database.User `json:"user"`
}

// CreateUserRequest representa os dados para criar um novo usuário.
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // Ex.: "superadmin", "administrator", "supervisor", "technician", "operator", "maintenance"
}

// UpdateUserRequest representa os dados para atualizar um usuário.
type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

// generateToken gera um token JWT para o usuário.
func generateToken(user database.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// AuthMiddleware verifica se o token JWT é válido e insere os claims no contexto.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userContextKey, claims)
		next(w, r.WithContext(ctx))
	}
}

// RequireRole garante que o usuário tenha uma das roles permitidas.
func RequireRole(roles ...string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(userContextKey).(*Claims)
			if !ok {
				http.Error(w, "Não autorizado", http.StatusUnauthorized)
				return
			}
			hasRole := false
			for _, role := range roles {
				if claims.Role == role {
					hasRole = true
					break
				}
			}
			if !hasRole {
				http.Error(w, "Permissão negada", http.StatusForbidden)
				return
			}
			next(w, r)
		}
	}
}

// LoginHandler processa o login do usuário.
func LoginHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Dados inválidos", http.StatusBadRequest)
			return
		}
		user, err := db.GetUserByUsername(req.Username)
		if err != nil {
			http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, "Senha incorreta", http.StatusUnauthorized)
			return
		}
		token, err := generateToken(user)
		if err != nil {
			http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
			return
		}
		// Remove a senha antes de retornar a resposta.
		user.Password = ""
		resp := LoginResponse{
			Token: token,
			User:  user,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

// CreateUserHandler cria um novo usuário.
func CreateUserHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Dados inválidos", http.StatusBadRequest)
			return
		}
		allowedRoles := map[string]bool{
			"superadmin":    true,
			"administrator": true,
			"supervisor":    true,
			"technician":    true,
			"operator":      true,
			"maintenance":   true,
		}
		if !allowedRoles[req.Role] {
			http.Error(w, "Role inválida", http.StatusBadRequest)
			return
		}
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Erro ao processar senha", http.StatusInternalServerError)
			return
		}
		now := time.Now()
		user := database.User{
			Username:  req.Username,
			Password:  string(hashedPass),
			Role:      req.Role,
			CreatedAt: now,
			UpdatedAt: now,
		}
		createdUser, err := db.CreateUser(user)
		if err != nil {
			http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
			return
		}
		createdUser.Password = ""
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createdUser)
	}
}

// UpdateUserHandler atualiza um usuário existente.
func UpdateUserHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		var req UpdateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Dados inválidos", http.StatusBadRequest)
			return
		}
		user, err := db.GetUserByID(id)
		if err != nil {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}
		// Atualiza os campos conforme enviados
		if req.Username != "" {
			user.Username = req.Username
		}
		if req.Role != "" {
			user.Role = req.Role
		}
		if req.Password != "" {
			hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "Erro ao processar senha", http.StatusInternalServerError)
				return
			}
			user.Password = string(hashedPass)
		}
		user.UpdatedAt = time.Now()
		if err := db.UpdateUser(user); err != nil {
			http.Error(w, "Erro ao atualizar usuário", http.StatusInternalServerError)
			return
		}
		user.Password = ""
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// DeleteUserHandler remove um usuário existente.
func DeleteUserHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		if err := db.DeleteUser(id); err != nil {
			http.Error(w, "Erro ao deletar usuário", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
