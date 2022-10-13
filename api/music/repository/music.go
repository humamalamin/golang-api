package repository

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	musicDomainEntity "test-fullstack/api/music/domain/entity"
	musicDomainInterface "test-fullstack/api/music/domain/interface"
)

type music struct {
	DB *gorm.DB
}

func NewMusicRepository(database *gorm.DB) musicDomainInterface.MusicRepository {
	repo := new(music)
	repo.DB = database

	return repo
}

func (repo *music) Get(ctx context.Context) ([]musicDomainEntity.Music, error) {
	results := []musicDomainEntity.Music{}

	err := repo.DB.Find(&results).Error
	if err != nil {
		code := "[Repository] Get - 1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	return results, nil
}

func (repo *music) GetByID(ctx context.Context, uuid string) (*musicDomainEntity.Music, error) {
	result := musicDomainEntity.Music{}

	err := repo.DB.Where("uuid = ?", uuid).Find(&result).Error
	if err != nil {
		code := "[Repository] GetByID - 1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	return &result, nil
}

func (repo *music) Store(ctx context.Context, req musicDomainEntity.Music) error {
	err := repo.DB.Create(&req).Error
	if err != nil {
		code := "[Repository] Store - 1"
		log.Error().Err(err).Msg(code)
		return err
	}

	return nil
}

func (repo *music) Update(ctx context.Context, req musicDomainEntity.Music) error {
	update := map[string]interface{}{
		"updated_at":   time.Now(),
		"name":         req.Name,
		"album":        req.Album,
		"album_art":    req.AlbumArt,
		"singer":       req.Singer,
		"publish_date": req.PublishDate,
	}

	err := repo.DB.Model(musicDomainEntity.Music{}).Where("uuid = ?", req.UUID).Updates(update).Error
	if err != nil {
		code := "[Repository] Update - 1"
		log.Error().Err(err).Msg(code)
		return err
	}

	return nil
}

func (repo *music) Delete(ctx context.Context, uuid string) error {
	err := repo.DB.Where("uuid = ?", uuid).Delete(&musicDomainEntity.Music{}).Error
	if err != nil {
		code := "[Repository] Delete - 1"
		log.Error().Err(err).Msg(code)
		return err
	}

	return nil
}
