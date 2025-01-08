package v1

import (
	"database/sql"
	"homework/specs"
)

// Быстрая проверка актуальности текущего интерфейса сервера.
var _ specs.ServerInterface = &apiServer{}

type apiServer struct {
	DB *sql.DB
}

func NewAPIServer(db *sql.DB) specs.ServerInterface {
	return &apiServer{db}
}
