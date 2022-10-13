package usecase

import (
	"context"
	musicDomainEntity "test-fullstack/api/music/domain/entity"
	musicDomainInterface "test-fullstack/api/music/domain/interface"
	musicRepository "test-fullstack/api/music/repository"
	"test-fullstack/packages/config"
	"test-fullstack/packages/manager"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Music struct {
	repo musicDomainInterface.MusicRepository
	cfg  *config.Config
}

func NewMusicUsecase(mgr manager.Manager) musicDomainInterface.MusicUsecase {
	usecase := new(Music)
	usecase.repo = musicRepository.NewMusicRepository(mgr.GetGorm())
	usecase.cfg = mgr.GetConfig()

	return usecase
}

func (uc *Music) Get(ctx context.Context) ([]musicDomainEntity.Music, error) {
	results, err := uc.repo.Get(ctx)
	if err != nil {
		code := "[Usecase] Get - 1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	return results, nil
}

func (uc *Music) GetByID(ctx context.Context, uuid string) (*musicDomainEntity.Music, error) {
	result, err := uc.repo.GetByID(ctx, uuid)
	if err != nil {
		code := "[Usecase] GetByID - 1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	return result, nil
}

func (uc *Music) Store(ctx context.Context, req musicDomainEntity.RequestMusic) error {
	updatedAt := time.Now()
	publishDate, _ := time.Parse("2006-01-02 15:04:05", *req.PublishDate)
	requestBody := musicDomainEntity.Music{
		UUID:        uuid.New(),
		Name:        req.Name,
		Album:       req.Album,
		AlbumArt:    req.AlbumArt,
		Singer:      req.Singer,
		PublishDate: &publishDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   &updatedAt,
	}

	err := uc.repo.Store(ctx, requestBody)
	if err != nil {
		code := "[Usecase] Store - 1"
		log.Error().Err(err).Msg(code)
		return err
	}

	return nil
}

func (uc *Music) Update(ctx context.Context, req musicDomainEntity.RequestMusic, id string) error {
	updatedAt := time.Now()
	publishDate, _ := time.Parse("2006-01-02 15:04:05", *req.PublishDate)
	requestBody := musicDomainEntity.Music{
		UUID:        uuid.MustParse(id),
		Name:        req.Name,
		Album:       req.Album,
		AlbumArt:    req.AlbumArt,
		Singer:      req.Singer,
		PublishDate: &publishDate,
		UpdatedAt:   &updatedAt,
	}

	err := uc.repo.Update(ctx, requestBody)
	if err != nil {
		code := "[Usecase] Update - 1"
		log.Error().Err(err).Msg(code)
		return err
	}

	return nil
}

func (uc *Music) Delete(ctx context.Context, uuid string) error {
	err := uc.repo.Delete(ctx, uuid)
	if err != nil {
		code := "[Usecase] Delete - 1"
		log.Error().Err(err).Msg(code)
		return err
	}

	return nil
}
