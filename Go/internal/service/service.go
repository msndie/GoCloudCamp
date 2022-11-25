package service

import "cmd/internal/repository"

type DistributedConfigService struct {
	repository *repository.ConfigRepository
}

func NewDistributedConfigService(r *repository.ConfigRepository) *DistributedConfigService {
	return &DistributedConfigService{r}
}
