package serv

import (
	openapi "github.com/iaPlotnikovv/New-breath/generated/go"
	"github.com/jmoiron/sqlx"
)

type BannerService struct {
	db *sqlx.DB
	openapi.DefaultAPIServicer
}

func NewBannerService(db *sqlx.DB) *BannerService {
	return &BannerService{
		db: db,
	}
}
