package database

import (
	"database/sql"
	"fmt"
	"time"
)

// LogLevel define os níveis de log.
type LogLevel string

const (
	LevelInfo  LogLevel = "INFO"
	LevelWarn  LogLevel = "WARN"
	LevelError LogLevel = "ERROR"
)

// Logger encapsula a conexão com o banco para gravar logs.
// Nesta implementação, os logs são inseridos em uma tabela "logs". Certifique-se de criar a tabela no seu banco.
type Logger struct {
	db *sql.DB
}

// NewLogger cria um novo logger utilizando a conexão do DB.
func NewLogger(db *DB) *Logger {
	return &Logger{db: db.conn}
}

// Log insere um log na tabela "logs".
func (l *Logger) Log(level LogLevel, message string, additional string) error {
	query := `INSERT INTO logs (timestamp, level, message, additional) VALUES (?, ?, ?, ?)`
	_, err := l.db.Exec(query, time.Now(), string(level), message, additional)
	if err != nil {
		return fmt.Errorf("erro ao inserir log: %v", err)
	}
	return nil
}

// Info registra um log de informação.
func (l *Logger) Info(message string, additional string) {
	if err := l.Log(LevelInfo, message, additional); err != nil {
		fmt.Printf("Logger Info erro: %v\n", err)
	}
}

// Warn registra um log de aviso.
func (l *Logger) Warn(message string, additional string) {
	if err := l.Log(LevelWarn, message, additional); err != nil {
		fmt.Printf("Logger Warn erro: %v\n", err)
	}
}

// Error registra um log de erro.
func (l *Logger) Error(message string, additional string) {
	if err := l.Log(LevelError, message, additional); err != nil {
		fmt.Printf("Logger Error erro: %v\n", err)
	}
}
