package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Port      string
	Cfg       mysql.Config
	SecretKey []byte
)

func LoadEnv() {
	// Caminho ABSOLUTO do arquivo config.go (onde esta função está)
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Não foi possível obter o caminho do arquivo de config")
	}

	// Pasta /config
	configDir := filepath.Dir(filePath)

	// .env dentro de /config
	envPath := filepath.Join(configDir, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Erro ao carregar .env em %s: %s", envPath, err)
	}
	log.Println("Variáveis de ambiente carregadas com sucesso")

	Port = os.Getenv("PORT")
	Cfg = mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_ADDR"),
		DBName:               os.Getenv("DB_DATABASE"),
		AllowNativePasswords: true,
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
