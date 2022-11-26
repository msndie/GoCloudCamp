package service

import (
	"cmd/internal/repository"
	ConfigService "cmd/proto"
	"log"
)

type DistributedConfigService struct {
	repository *repository.ConfigRepository
}

func NewDistributedConfigService(r *repository.ConfigRepository) *DistributedConfigService {
	return &DistributedConfigService{r}
}

func (d *DistributedConfigService) AddConfig(config *ConfigService.Config) bool {
	if err := d.repository.InsertInCollection(config); err != nil {
		log.Print(err.Error())
		return false
	}
	return true
}

func (d *DistributedConfigService) FindConfig(name string) *ConfigService.Config {
	config, err := d.repository.FindLastVersionInCollection(name)
	if err != nil {
		log.Print(err.Error())
		return nil
	}
	return config
}
