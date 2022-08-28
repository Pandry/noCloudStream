package api

import (
	"auther/models"
	"auther/repository"
)

type Backpack struct {
	repository.Repository
	Signer  models.JwtSigner
	Logger  models.Logger
	Sigterm *models.Sigterm
}

func NewBackpack(repo repository.Repository, signer models.JwtSigner, logger models.Logger, sigterm *models.Sigterm) Backpack {
	return Backpack{repo, signer, logger, sigterm}
}
