package musicDomainEntity

import (
	"time"

	"github.com/google/uuid"
)

type Music struct {
	UUID        uuid.UUID  `json:"uuid"`
	Name        string     `json:"name"`
	Album       string     `json:"album"`
	AlbumArt    *string    `json:"album_art"`
	Singer      string     `json:"singer"`
	PublishDate *time.Time `json:"publish_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type RequestMusic struct {
	Name        string  `json:"name"`
	Album       string  `json:"album"`
	AlbumArt    *string `json:"album_art"`
	Singer      string  `json:"singer"`
	PublishDate *string `json:"publish_date"`
}
