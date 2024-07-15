package postgresql

import (
	"fmt"
	"log/slog"
	"time"

	"cinematic.com/sso/internal/infrastructure/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	config *config.StorageConfig
	logger *slog.Logger
	DB     *sqlx.DB
}

func (p *Storage) Connect() error {
	dsn := p.makeDsn()

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		p.logger.Error("Failed to connect to db", slog.Any("error", err))
		return err
	}

	time.Sleep(time.Second * 2)

	if err := db.Ping(); err != nil {
		p.logger.Error("Failed to ping db", slog.Any("error", err))
		return err
	}

	p.DB = db
	return nil
}

func (p *Storage) makeDsn() string {
	res := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", p.config.User, p.config.Password, p.config.Database, p.config.Host, p.config.Port)

	return res
}

func NewPostgreSQLStorage(
	config *config.StorageConfig,
	logger *slog.Logger,
) *Storage {
	return &Storage{
		config: config,
		logger: logger,
	}
}
