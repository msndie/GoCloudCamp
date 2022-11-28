package service

import (
	"cmd/internal/repository"
	pb "cmd/proto"
)

func checkEquality(lhs *pb.Config, rhs *pb.Config) bool {
	l1, l2 := len(lhs.Data), len(rhs.Data)

	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		if lhs.Data[i].Key != rhs.Data[i].Key ||
			lhs.Data[i].Value != rhs.Data[i].Value {
			return false
		}
	}
	return true
}

type DistributedConfigService struct {
	repository *repository.ConfigRepository
}

func NewDistributedConfigService(r *repository.ConfigRepository) *DistributedConfigService {
	return &DistributedConfigService{repository: r}
}

func (d *DistributedConfigService) AddConfig(config *pb.Config) (bool, error) {
	last, err := d.repository.FindLastVersionInCollection(config.Service)
	if err != nil || (last != nil && checkEquality(last, config)) {
		return false, err
	}
	b, err := d.repository.InsertInCollection(config)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (d *DistributedConfigService) FindConfig(name string) (*pb.Config, error) {
	config, err := d.repository.FindLastVersionInCollection(name)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (d *DistributedConfigService) GetAllVersionsOfConfig(name string) ([]*pb.Config, error) {
	configs, err := d.repository.FindAllInCollection(name)
	if err != nil {
		return nil, err
	}
	return configs, nil
}

func (d *DistributedConfigService) GetAllConfigs() ([]*pb.Config, error) {
	configs, err := d.repository.FindAllLastEntities()
	if err != nil {
		return nil, err
	}
	return configs, nil
}

func (d *DistributedConfigService) DeleteConfig(name string) (bool, error) {
	err := d.repository.DeleteCollection(name)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (d *DistributedConfigService) IsExistsByName(name string) (bool, error) {
	return d.repository.CollectionExists(name)
}
