package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DBConfig contém as configurações de conexão.
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// DB representa a conexão com o banco.
type DB struct {
	conn *sql.DB
}

// PLC representa um PLC.
type PLC struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	IPAddress  string    `json:"ip_address"`
	Rack       int       `json:"rack"`
	Slot       int       `json:"slot"`
	Active     bool      `json:"active"`
	Status     string    `json:"status"`
	LastUpdate time.Time `json:"last_update"`
	TagCount   int       `json:"tag_count"` // Novo campo para quantidade de tags
}

// Tag representa uma tag associada a um PLC.
type Tag struct {
	ID             int    `json:"id"`
	PLCID          int    `json:"plc_id"`
	Name           string `json:"name"`
	DBNumber       int    `json:"db_number"`
	ByteOffset     int    `json:"byte_offset"`
	DataType       string `json:"data_type"`
	CanWrite       bool   `json:"can_write"`
	ScanRate       int    `json:"scan_rate"`
	MonitorChanges bool   `json:"monitor_changes"`
	Active         bool   `json:"active"`
}

// PLCStatus é usado para atualizar os campos status e last_update de um PLC.
type PLCStatus struct {
	PLCID      int
	Status     string
	LastUpdate time.Time
}

// NewDB cria uma nova conexão com o banco.
// O DSN inclui "parseTime=true" para converter automaticamente os campos DATETIME para time.Time.
func NewDB(conf DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return &DB{conn: conn}, nil
}

// Close fecha a conexão com o banco.
func (db *DB) Close() error {
	return db.conn.Close()
}

// GetActivePLCs retorna os PLCs ativos e conta as tags ativas associadas a cada um.
func (db *DB) GetActivePLCs() ([]PLC, error) {
	// A query faz um subselect para contar as tags ativas para cada PLC.
	query := `
    SELECT p.id, p.name, p.ip_address, p.rack, p.slot, p.active, p.status, p.last_update,
           (SELECT COUNT(*) FROM tags t WHERE t.plc_id = p.id AND t.active = true) AS tag_count
    FROM plcs p
    WHERE p.active = true`
	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plcs []PLC
	for rows.Next() {
		var plc PLC
		var lastUpdate sql.NullTime
		var status sql.NullString
		var tagCount sql.NullInt64
		if err := rows.Scan(&plc.ID, &plc.Name, &plc.IPAddress, &plc.Rack, &plc.Slot, &plc.Active, &status, &lastUpdate, &tagCount); err != nil {
			return nil, err
		}
		if status.Valid {
			plc.Status = status.String
		}
		if lastUpdate.Valid {
			plc.LastUpdate = lastUpdate.Time
		}
		if tagCount.Valid {
			plc.TagCount = int(tagCount.Int64)
		} else {
			plc.TagCount = 0
		}
		plcs = append(plcs, plc)
	}
	return plcs, nil
}

// GetPLCByID retorna os dados de um PLC pelo ID.
func (db *DB) GetPLCByID(id int) (PLC, error) {
	var plc PLC
	var lastUpdate sql.NullTime
	var status sql.NullString
	query := "SELECT id, name, ip_address, rack, slot, active, status, last_update FROM plcs WHERE id = ?"
	err := db.conn.QueryRow(query, id).Scan(&plc.ID, &plc.Name, &plc.IPAddress, &plc.Rack, &plc.Slot, &plc.Active, &status, &lastUpdate)
	if err != nil {
		return plc, err
	}
	if status.Valid {
		plc.Status = status.String
	}
	if lastUpdate.Valid {
		plc.LastUpdate = lastUpdate.Time
	}
	// Opcional: buscar a quantidade de tags para este PLC.
	tagCountQuery := "SELECT COUNT(*) FROM tags WHERE plc_id = ? AND active = true"
	var count int
	if err := db.conn.QueryRow(tagCountQuery, id).Scan(&count); err == nil {
		plc.TagCount = count
	} else {
		plc.TagCount = 0
	}
	return plc, nil
}

// CreatePLC insere um novo PLC e retorna o PLC criado.
func (db *DB) CreatePLC(plc PLC) (PLC, error) {
	query := "INSERT INTO plcs (name, ip_address, rack, slot, active, status, last_update) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := db.conn.Exec(query, plc.Name, plc.IPAddress, plc.Rack, plc.Slot, plc.Active, plc.Status, plc.LastUpdate)
	if err != nil {
		return plc, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return plc, err
	}
	plc.ID = int(id)
	return plc, nil
}

// UpdatePLC atualiza os dados de um PLC.
func (db *DB) UpdatePLC(plc PLC) error {
	query := "UPDATE plcs SET name = ?, ip_address = ?, rack = ?, slot = ?, active = ?, status = ?, last_update = ? WHERE id = ?"
	_, err := db.conn.Exec(query, plc.Name, plc.IPAddress, plc.Rack, plc.Slot, plc.Active, plc.Status, plc.LastUpdate, plc.ID)
	return err
}

// DeletePLC remove um PLC e suas tags associadas utilizando transação.
func (db *DB) DeletePLC(id int) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	// Deleta as tags associadas ao PLC
	queryTags := "DELETE FROM tags WHERE plc_id = ?"
	if _, err := tx.Exec(queryTags, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao deletar tags associadas: %w", err)
	}

	// Deleta o PLC
	queryPLC := "DELETE FROM plcs WHERE id = ?"
	if _, err := tx.Exec(queryPLC, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao deletar PLC: %w", err)
	}

	return tx.Commit()
}

// GetPLCTags retorna as tags associadas a um PLC.
func (db *DB) GetPLCTags(plcID int) ([]Tag, error) {
	query := "SELECT id, name, db_number, byte_offset, data_type, can_write, scan_rate, monitor_changes, active FROM tags WHERE plc_id = ? AND active = true"
	rows, err := db.conn.Query(query, plcID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var tag Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.DBNumber, &tag.ByteOffset, &tag.DataType, &tag.CanWrite, &tag.ScanRate, &tag.MonitorChanges, &tag.Active); err != nil {
			return nil, err
		}
		tag.PLCID = plcID
		tags = append(tags, tag)
	}
	return tags, nil
}

// CreateTag insere uma nova tag para um PLC.
func (db *DB) CreateTag(tag Tag) (Tag, error) {
	query := "INSERT INTO tags (plc_id, name, db_number, byte_offset, data_type, can_write, scan_rate, monitor_changes, active) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := db.conn.Exec(query, tag.PLCID, tag.Name, tag.DBNumber, tag.ByteOffset, tag.DataType, tag.CanWrite, tag.ScanRate, tag.MonitorChanges, tag.Active)
	if err != nil {
		return tag, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return tag, err
	}
	tag.ID = int(id)
	return tag, nil
}

// UpdateTag atualiza os dados de uma tag.
func (db *DB) UpdateTag(tag Tag) error {
	query := "UPDATE tags SET name = ?, db_number = ?, byte_offset = ?, data_type = ?, can_write = ?, scan_rate = ?, monitor_changes = ?, active = ? WHERE id = ?"
	_, err := db.conn.Exec(query, tag.Name, tag.DBNumber, tag.ByteOffset, tag.DataType, tag.CanWrite, tag.ScanRate, tag.MonitorChanges, tag.Active, tag.ID)
	return err
}

// DeleteTag remove uma tag.
func (db *DB) DeleteTag(tagID int) error {
	query := "DELETE FROM tags WHERE id = ?"
	_, err := db.conn.Exec(query, tagID)
	return err
}

// UpdatePLCStatus atualiza os campos status e last_update de um PLC.
func (db *DB) UpdatePLCStatus(status PLCStatus) error {
	query := "UPDATE plcs SET status = ?, last_update = ? WHERE id = ?"
	result, err := db.conn.Exec(query, status.Status, status.LastUpdate, status.PLCID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		fmt.Printf("Nenhuma linha atualizada para o PLC ID %d (valores possivelmente inalterados)\n", status.PLCID)
		return nil
	}
	return nil
}
