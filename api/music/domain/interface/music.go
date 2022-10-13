package musicDomainInterface

import (
	"context"
	"net/http"

	musicDomainEntity "test-fullstack/api/music/domain/entity"
)

type MusicHandler interface {
	Get() http.Handler
	GetByID() http.Handler
	Store() http.Handler
	Update() http.Handler
	Delete() http.Handler
}

type MusicRepository interface {
	Get(ctx context.Context) ([]musicDomainEntity.Music, error)
	GetByID(ctx context.Context, uuid string) (*musicDomainEntity.Music, error)
	Store(ctx context.Context, req musicDomainEntity.Music) error
	Update(ctx context.Context, req musicDomainEntity.Music) error
	Delete(ctx context.Context, uuid string) error
}

type MusicUsecase interface {
	Get(ctx context.Context) ([]musicDomainEntity.Music, error)
	GetByID(ctx context.Context, uuid string) (*musicDomainEntity.Music, error)
	Store(ctx context.Context, req musicDomainEntity.RequestMusic) error
	Update(ctx context.Context, req musicDomainEntity.RequestMusic, uuid string) error
	Delete(ctx context.Context, uuid string) error
}
