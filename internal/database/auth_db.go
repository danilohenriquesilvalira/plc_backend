package database

import (
	"database/sql"
	"fmt"
	"time"
)

// User representa um usuário do sistema
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // O "-" esconde a senha do JSON
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUser cria um novo usuário no banco de dados
func (db *DB) CreateUser(user User) (User, error) {
	query := `
		INSERT INTO users (username, password, role, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`
	now := time.Now()
	result, err := db.conn.Exec(query,
		user.Username,
		user.Password,
		user.Role,
		now,
		now,
	)
	if err != nil {
		return User{}, fmt.Errorf("erro ao criar usuário: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return User{}, fmt.Errorf("erro ao obter ID do usuário: %v", err)
	}

	user.ID = int(id)
	user.CreatedAt = now
	user.UpdatedAt = now
	return user, nil
}

// GetUserByUsername busca um usuário pelo nome de usuário
func (db *DB) GetUserByUsername(username string) (User, error) {
	var user User
	query := `
		SELECT id, username, password, role, created_at, updated_at
		FROM users
		WHERE username = ?
	`
	err := db.conn.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("usuário não encontrado")
		}
		return User{}, fmt.Errorf("erro ao buscar usuário: %v", err)
	}
	return user, nil
}

// GetUserByID busca um usuário pelo ID
func (db *DB) GetUserByID(id int) (User, error) {
	var user User
	query := `
		SELECT id, username, password, role, created_at, updated_at
		FROM users
		WHERE id = ?
	`
	err := db.conn.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("usuário não encontrado")
		}
		return User{}, fmt.Errorf("erro ao buscar usuário: %v", err)
	}
	return user, nil
}

// UpdateUser atualiza os dados de um usuário
func (db *DB) UpdateUser(user User) error {
	query := `
		UPDATE users
		SET username = ?,
			role = ?,
			updated_at = ?
		WHERE id = ?
	`
	now := time.Now()
	result, err := db.conn.Exec(query,
		user.Username,
		user.Role,
		now,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar atualização: %v", err)
	}
	if rows == 0 {
		return fmt.Errorf("usuário não encontrado")
	}

	return nil
}

// UpdateUserPassword atualiza a senha de um usuário
func (db *DB) UpdateUserPassword(userID int, hashedPassword string) error {
	query := `
		UPDATE users
		SET password = ?,
			updated_at = ?
		WHERE id = ?
	`
	result, err := db.conn.Exec(query, hashedPassword, time.Now(), userID)
	if err != nil {
		return fmt.Errorf("erro ao atualizar senha: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar atualização: %v", err)
	}
	if rows == 0 {
		return fmt.Errorf("usuário não encontrado")
	}

	return nil
}

// DeleteUser remove um usuário
func (db *DB) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := db.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar usuário: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar deleção: %v", err)
	}
	if rows == 0 {
		return fmt.Errorf("usuário não encontrado")
	}

	return nil
}

// ListUsers retorna todos os usuários
func (db *DB) ListUsers() ([]User, error) {
	query := `
		SELECT id, username, role, created_at, updated_at
		FROM users
		ORDER BY username
	`
	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar usuários: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler usuário: %v", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao percorrer usuários: %v", err)
	}

	return users, nil
}
