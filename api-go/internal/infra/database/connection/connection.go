package connection

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/Gabriel-Schiestl/api-go/internal/config"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConfig *config.DbConfig
var Db *gorm.DB

func SetupConfig(host, user, password, port, name string) *sql.DB {
	dbPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Error converting DB_PORT to int")
	}

	DbConfig = config.NewDbConfig(host, user, password, name, dbPort)

	Db, err = gorm.Open(postgres.Open(DbConfig.ToString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	sqlDb, err := Db.DB()
	if err != nil {
		log.Fatalf("Error getting DB connection: %v", err)
	}

	Db.AutoMigrate(entities.Event{})
	
	// Forçar migração da tabela users
	if err := Db.AutoMigrate(&entities.User{}); err != nil {
		log.Printf("Warning: Failed to migrate User table: %v", err)
	}
	
	// Verificar se as colunas existem e adicionar se necessário
	if !Db.Migrator().HasColumn(&entities.User{}, "password") {
		if err := Db.Migrator().AddColumn(&entities.User{}, "password"); err != nil {
			log.Printf("Warning: Failed to add password column: %v", err)
		}
	}
	
	if !Db.Migrator().HasColumn(&entities.User{}, "user_type") {
		if err := Db.Migrator().AddColumn(&entities.User{}, "user_type"); err != nil {
			log.Printf("Warning: Failed to add user_type column: %v", err)
		}
	}
	
	Db.AutoMigrate(entities.Auth{})

	return sqlDb
}