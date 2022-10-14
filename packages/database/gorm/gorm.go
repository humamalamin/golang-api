package gormDatabase

import (
	"test-fullstack/packages/config"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm interface {
	Connect() (*gorm.DB, error)
}

type Options struct {
	master  string
	maxOpen int
	maxIdle int
}

func NewGorm(cfg *config.Config) Gorm {
	opt := new(Options)
	opt.master = cfg.DatabaseDriver + "://" + " user=" + cfg.DatabaseUsername + " password=" + cfg.DatabasePassword + " dbname=" + cfg.DatabaseName + " port=" + cfg.DatabasePort + " sslmode=" + cfg.DatabaseSsl + " TimeZone=" + cfg.AppTz
	opt.maxIdle = cfg.DatabaseMaxIdleConnections
	opt.maxOpen = cfg.DatabaseMaxOpenConnections

	return opt
}

func (o *Options) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(o.master), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("[Connect-1 Failed To Connect Gorm")
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(o.maxOpen)
	sqlDB.SetMaxIdleConns(o.maxIdle)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
