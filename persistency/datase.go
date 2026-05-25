package persistency

import (
	"database/sql"
	"log"
)

func Connect() (*sql.DB, error) {
	// TODO: Adicionar importação de config ou usar variáveis de ambiente
	// Por enquanto usando string de conexão padrão
	dsn := "user:password@tcp(localhost:3306)/treehousedb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Erro ao abrir conexão com o DB")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Println("Erro ao pingar no DB")
		return nil, err
	}

	return db, nil
}
