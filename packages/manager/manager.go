package manager

import (
	"database/sql"
	"os"
	paginationHelper "test-fullstack/helpers/pagination"
	gocacheCache "test-fullstack/packages/cache/go-cache"
	"test-fullstack/packages/config"
	gormDatabase "test-fullstack/packages/database/gorm"
	"test-fullstack/packages/server"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Manager interface {
	GetConfig() *config.Config
	GetServer() *server.Server
	GetMysql() *sql.DB
	GetGorm() *gorm.DB
	GetPagination() paginationHelper.Pagination
	GetGoCache() gocacheCache.GoCache
}

type manager struct {
	config       *config.Config
	server       *server.Server
	dbMysql      *sql.DB
	dbGorm       *gorm.DB
	pagination   paginationHelper.Pagination
	goCacheCache gocacheCache.GoCache
}

func NewInit() (Manager, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Error().Err(err).Msg("[NewInit-1] Failed to Initialize Configuration")
		return nil, err
	}

	srv := server.NewServer(cfg)

	dbGorm, err := gormDatabase.NewGorm(cfg).Connect()
	if err != nil {
		log.Error().Err(err).Msg("[NewInit-3] Failed to Initialize Database Gorm")
		return nil, err
	}

	log.Logger = log.With().Caller().Logger()
	if cfg.AppIsDev {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}).With().Caller().Logger()
	}

	paginationHelper := paginationHelper.NewPagination()

	goCache := gocacheCache.NewGoCache()

	return &manager{
		config:       cfg,
		server:       srv,
		dbGorm:       dbGorm,
		pagination:   paginationHelper,
		goCacheCache: goCache,
	}, nil
}

func (sm *manager) GetConfig() *config.Config {
	return sm.config
}

func (sm *manager) GetServer() *server.Server {
	return sm.server
}

func (sm *manager) GetMysql() *sql.DB {
	return sm.dbMysql
}

func (sm *manager) GetGorm() *gorm.DB {
	return sm.dbGorm
}

func (sm *manager) GetPagination() paginationHelper.Pagination {
	return sm.pagination
}

func (sm *manager) GetGoCache() gocacheCache.GoCache {
	return sm.goCacheCache
}
